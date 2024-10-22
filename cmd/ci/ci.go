package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/ko/pkg/build"
	"github.com/google/ko/pkg/publish"
	"github.com/verifa/website/pkg/watcher"
)

const (
	goCILint = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	goFumpt  = "mvdan.cc/gofumpt"
	goAir    = "github.com/cosmtrek/air"
	goTempl  = "github.com/a-h/templ/cmd/templ"
)

const (
	baseImage              = "cgr.dev/chainguard/static:latest"
	targetRepo             = "europe-north1-docker.pkg.dev/verifa-website/website"
	targetImage            = targetRepo + "/website"
	importpath             = "github.com/verifa/website/cmd/website"
	cloudRunServiceProd    = "prod-website-service"
	cloudRunServiceStaging = "staging-website-service"

	region = "europe-north1"
)

const (
	curDir = "."
	recDir = "./..."
)

var gitCommit = "dev"

func main() {
	var dev, build, lint, test, preview, pr bool
	var deploy string
	flag.BoolVar(&dev, "dev", false, "run the website locally")
	flag.BoolVar(&build, "build", false, "build the website locally")
	flag.BoolVar(&lint, "lint", false, "lint the code")
	flag.BoolVar(&test, "test", false, "run the tests")
	flag.StringVar(
		&deploy,
		"deploy",
		"",
		"deploy the website to this env (staging or prod)",
	)
	flag.BoolVar(&preview, "preview", false, "run a local preview environment")
	flag.BoolVar(&pr, "pr", false, "run the pull request checks")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Handle recover and cancel context.
	defer func() {
		if err := recover(); err != nil {
			cancel()
			log.Fatal("panic occurred:", err)
		}
	}()

	// Handle signals and create cancel context.
	ctx, stop := signal.NotifyContext(
		ctx,
		os.Interrupt,
		os.Kill,
	)
	defer stop()

	if dev {
		Dev(ctx)
	}
	if build {
		_ = KoBuild(ctx, WithKoLocal())
	}
	if preview {
		Preview(ctx)
	}
	if lint {
		Lint(ctx)
	}
	if test {
		Test(ctx)
	}
	if pr {
		PullRequest(ctx)
		HasGitDiff(ctx)
	}
	if deploy != "" {
		Deploy(ctx, deploy)
	}
}

func Dev(ctx context.Context) {
	fmt.Println("🚀 starting dev server")
	Watch(ctx)
	<-ctx.Done()
}

func Watch(ctx context.Context) {
	fmt.Println("👀 watching for changes")

	runner := Runner{}
	if err := watcher.WatchFilesystem(ctx, watcher.WatchOptions{
		RunOnStart:    true,
		Name:          "go",
		Dir:           ".",
		IncludeFiles:  []string{"app.css"},
		IncludeSuffix: []string{".go", ".md", ".templ", ".bib"},
		ExcludeSuffix: []string{"_templ.go", "_test.go"},
		Batch:         200 * time.Millisecond,
		Fn: func(paths []string) {
			fmt.Println("📝 source file changed: ", paths)
			if err := Generate(ctx); err != nil {
				fmt.Printf("⚠️ generate: %s\n", err.Error())
			}
			fmt.Println("🚀 restarting website")
			if err := runner.Stop(); err != nil {
				fmt.Printf("❌ stopping website: %s\n", err)
				return
			}
			if err := runner.Start(ctx); err != nil {
				fmt.Printf("❌ starting website: %s\n", err)
				return
			}
		},
	}); err != nil {
		panic(fmt.Sprintf("watching filesystem: %s", err))
	}
}

func Generate(ctx context.Context) error {
	fmt.Println("📝 generating content")
	wg := sync.WaitGroup{}
	wg.Add(2)
	var templErr error
	var tailwindErr error
	go func() {
		templErr = errorf("templ: %w", TemplGenerate(ctx))
		wg.Done()
	}()
	go func() {
		tailwindErr = errorf("tailwind: %w", TailwindGenerate(ctx))
		wg.Done()
	}()
	wg.Wait()
	if errs := errors.Join(templErr, tailwindErr); errs != nil {
		return errs
	}
	fmt.Println("✅ content generated")
	return nil
}

func TemplGenerate(ctx context.Context) error {
	args := []string{goTempl, "generate"}
	return GoRun(ctx, args...)
}

func TailwindGenerate(ctx context.Context) error {
	return NpxRun(
		ctx,
		"tailwindcss",
		"--config",
		"./tailwind.config.cjs",
		"--input",
		"./app.css",
		"--output",
		"./dist/tailwind.css",
		"--minify",
	)
}

type KoOption func(*koOptions)

func WithKoLocal() KoOption {
	return func(o *koOptions) {
		o.local = true
	}
}

type koOptions struct {
	local bool
}

func KoBuild(ctx context.Context, opts ...KoOption) string {
	fmt.Println("🏗️ building container image")
	opt := &koOptions{}
	for _, o := range opts {
		o(opt)
	}

	b, err := build.NewGo(
		ctx,
		"./cmd/website",
		build.WithPlatforms("linux/amd64"),
		build.WithBaseImages(
			func(ctx context.Context, _ string) (name.Reference, build.Result, error) {
				ref := name.MustParseReference(baseImage)
				base, err := remote.Index(ref, remote.WithContext(ctx))
				return ref, base, err
			},
		),
		build.WithConfig(map[string]build.Config{
			"github.com/verifa/website/cmd/website": {
				Ldflags: []string{
					"-X main.buildGitCommit=" + gitCommit,
				},
			},
		}),
	)
	if err != nil {
		log.Fatalf("ko: creating build interface: %v", err)
	}
	r, err := b.Build(ctx, importpath)
	if err != nil {
		log.Fatalf("ko: building: %v", err)
	}
	var (
		pub    publish.Interface
		pubErr error
	)
	if opt.local {
		pub, pubErr = publish.NewDaemon(
			targetRepoNamer,
			[]string{gitCommit},
		)
	} else {
		pub, pubErr = publish.NewDefault(targetRepo,
			publish.WithNamer(targetRepoNamer),
			publish.WithTags([]string{gitCommit}),
			publish.WithAuthFromKeychain(authn.DefaultKeychain))
	}
	if pubErr != nil {
		log.Fatalf("ko: creating publish interface: %v", err)
	}
	ref, err := pub.Publish(ctx, r, importpath)
	if err != nil {
		log.Fatalf("ko: publishing: %v", err)
	}
	fmt.Println(ref.String())
	fmt.Println("✅ container image published")
	return ref.String()
}

func targetRepoNamer(s1, s2 string) string {
	return targetImage
}

func Preview(ctx context.Context) {
	fmt.Println("🧪 starting preview")
	iferr(Generate(ctx))
	ref := KoBuild(ctx, WithKoLocal())
	iferr(DockerRun(
		ctx,
		"run",
		"--rm",
		"-i",
		"-p",
		"3000:3000",
		ref,
	))
}

func Lint(ctx context.Context) {
	fmt.Println("🧹 code linting")
	iferr(Generate(ctx))
	iferr(Go(ctx, "mod", "tidy"))
	iferr(Go(ctx, "mod", "verify"))
	iferr(GoRun(ctx, goFumpt, "-w", "-extra", curDir))
	iferr(GoRun(ctx, goCILint, "-v", "run", recDir))
	fmt.Println("✅ code linted")
}

func Test(ctx context.Context) {
	fmt.Println("🧪 running tests")
	iferr(Generate(ctx))
	iferr(Go(ctx, "test", "-v", recDir))
	fmt.Println("✅ tests passed")
}

func PullRequest(ctx context.Context) {
	iferr(Generate(ctx))
	Lint(ctx)
	Test(ctx)
	fmt.Println("✅ pull request checks passed")
}

// HasGitDiff displays the git diff and errors if there is a diff
func HasGitDiff(ctx context.Context) {
	cmd := exec.CommandContext(ctx, "git", "--no-pager", "diff")
	slog.Info("exec", slog.String("cmd", cmd.String()))
	b, err := cmd.CombinedOutput()
	iferr(err)
	if len(b) == 0 {
		return
	}
	buf := bytes.NewBuffer(b)
	fmt.Println("❌ git diff is not empty:")
	fmt.Println(buf.String())
	os.Exit(1)
}

func Deploy(ctx context.Context, deploy string) {
	var cloudRunService string
	switch deploy {
	case "staging":
		cloudRunService = cloudRunServiceStaging
	case "prod":
		cloudRunService = cloudRunServiceProd
	default:
		if !strings.HasPrefix(deploy, "pr-") {
			panic("invalid deploy env")
		}
		cloudRunService = deploy + "-website-service"
	}
	fmt.Println("🚢 deploying to", deploy)
	iferr(Generate(ctx))
	ref := KoBuild(ctx)
	iferr(GCloudRun(
		ctx,
		"run",
		"deploy",
		cloudRunService,
		"--image",
		ref,
		"--region",
		region,
	))
	fmt.Println("✅ deployed to", deploy)
}

func Go(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "go", args...)
	slog.Info("exec", slog.String("cmd", cmd.String()))
	defer slog.Info("done", slog.String("cmd", cmd.String()))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_ = os.Stderr.Sync()
		_ = os.Stdout.Sync()
		return fmt.Errorf("go: %s", err)
	}
	return nil
}

func GoRun(ctx context.Context, args ...string) error {
	return Go(ctx, append([]string{"run", "-mod=readonly"}, args...)...)
}

func NpxRun(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "npx", args...)
	slog.Info("exec", slog.String("cmd", cmd.String()))
	defer slog.Info("done", slog.String("cmd", cmd.String()))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_ = os.Stderr.Sync()
		_ = os.Stdout.Sync()
		return fmt.Errorf("npx: %w", err)
	}
	return nil
}

func DockerRun(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "docker", args...)
	slog.Info("exec", slog.String("cmd", cmd.String()))
	defer slog.Info("done", slog.String("cmd", cmd.String()))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_ = os.Stderr.Sync()
		_ = os.Stdout.Sync()
		return fmt.Errorf("docker: %s", err)
	}
	return nil
}

func GCloudRun(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "gcloud", args...)
	slog.Info("exec", slog.String("cmd", cmd.String()))
	defer slog.Info("done", slog.String("cmd", cmd.String()))
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_ = os.Stderr.Sync()
		_ = os.Stdout.Sync()
		return fmt.Errorf("docker: %s", err)
	}
	return nil
}

func iferr(err error) {
	if err != nil {
		panic(err)
	}
}

func errorf(format string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf(format, err)
}

func commitSHA() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	b, err := cmd.CombinedOutput()
	s := strings.TrimSpace(string(b))
	return s, err
}

func init() {
	commit, err := commitSHA()
	iferr(err)
	gitCommit = commit
}

type Runner struct {
	mu     sync.Mutex
	cmd    *exec.Cmd
	cancel context.CancelFunc
}

func (r *Runner) Start(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Build the website first.
	if err := r.build(ctx); err != nil {
		return fmt.Errorf("building website: %w", err)
	}
	ctx, cancel := context.WithCancel(ctx)
	cmd := exec.CommandContext(
		ctx,
		"./build/website",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		cancel()
		return fmt.Errorf("starting website: %w", err)
	}
	r.cmd = cmd
	r.cancel = cancel
	return nil
}

func (r *Runner) build(ctx context.Context) error {
	cmd := exec.CommandContext(
		ctx,
		"go",
		"build",
		"-o",
		"./build/website",
		"./cmd/website/website.go",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go build: %w", err)
	}
	return nil
}

func (r *Runner) Stop() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.cmd != nil {
		r.cancel()
		if err := r.cmd.Wait(); err != nil {
			var exitErr *exec.ExitError
			if !errors.As(err, &exitErr) {
				return fmt.Errorf("waiting for website to stop: %w", err)
			}
		}
	}
	r.cmd = nil
	return nil
}
