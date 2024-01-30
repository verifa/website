package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/ko/pkg/build"
	"github.com/google/ko/pkg/publish"
)

const (
	goCILint = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	goFumpt  = "mvdan.cc/gofumpt"
	goAir    = "github.com/cosmtrek/air"
	goTempl  = "github.com/a-h/templ/cmd/templ"
)

const (
	baseImage = "cgr.dev/chainguard/static:latest"
	// baseImage  = "gcr.io/distroless/static-debian11:nonroot"
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
	var dev, build, lint, test, run, preview, pr bool
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
	flag.BoolVar(&run, "run", false, "run the website locally")
	flag.BoolVar(&preview, "preview", false, "run a local preview environment")
	flag.BoolVar(&pr, "pr", false, "run the pull request checks")
	flag.Parse()

	// Handle signals and create cancel context.
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		os.Kill,
	)
	defer cancel()

	// Handle recover and cancel context.
	defer func() {
		if err := recover(); err != nil {
			cancel()
			log.Println("panic occurred:", err)
		}
	}()

	if lint {
		Lint(ctx)
	}
	if test {
		Test(ctx)
	}
	if dev {
		Dev(ctx)
	}
	if run {
		Run(ctx)
	}
	if build {
		_ = KoBuild(ctx, WithKoLocal())
	}
	if preview {
		Preview(ctx)
	}
	if pr {
		PullRequest(ctx)
	}
	if deploy != "" {
		Deploy(ctx, deploy)
	}
}

func Dev(ctx context.Context) {
	Watch(ctx)
	fmt.Println("🚀 starting dev server")
	iferr(GoRun(ctx, goAir, "-c", ".air.toml"))
}

func Run(ctx context.Context) {
	Generate(ctx)
	iferr(Go(ctx, "run", "-o", "./build/website", "./cmd/website/main.go"))
}

func Watch(ctx context.Context) {
	fmt.Println("👀 watching for changes")
	go func() {
		iferr(GoRun(ctx, goTempl, "generate", "--watch"))
	}()
	go func() {
		iferr(
			NpxRun(
				ctx,
				"tailwindcss",
				"build",
				"-i",
				"./src/app.css",
				"-o",
				"./dist/tailwind.css",
				"--minify",
				"--watch",
			),
		)
	}()
}

func Generate(ctx context.Context) {
	fmt.Println("📝 generating content")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		iferr(GoRun(ctx, goTempl, "generate"))
		wg.Done()
	}()
	go func() {
		iferr(
			NpxRun(
				ctx,
				"tailwindcss",
				"build",
				"-i",
				"./src/app.css",
				"-o",
				"./dist/tailwind.css",
				"--minify",
			),
		)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("✅ content generated")
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
	iferr(Go(ctx, "mod", "tidy"))
	iferr(Go(ctx, "mod", "verify"))
	iferr(GoRun(ctx, goFumpt, "-w", "-extra", curDir))
	iferr(GoRun(ctx, goCILint, "-v", "run", recDir))
	fmt.Println("✅ code linted")
}

func Test(ctx context.Context) {
	fmt.Println("🧪 running tests")
	iferr(Go(ctx, "test", "-v", recDir))
	fmt.Println("✅ tests passed")
}

func PullRequest(ctx context.Context) {
	Lint(ctx)
	Test(ctx)
	fmt.Println("✅ pull request checks passed")
}

func Deploy(ctx context.Context, deploy string) {
	var cloudRunService string
	switch deploy {
	case "staging":
		cloudRunService = cloudRunServiceStaging
	case "prod":
		cloudRunService = cloudRunServiceProd
	default:
		panic("invalid deploy env")
	}
	fmt.Println("🚢 deploying to", deploy)
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
		return fmt.Errorf("npx: %s", err)
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
