package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/golingon/lingon/pkg/terra"
	"github.com/golingon/lingon/pkg/x/sylt"
	google "github.com/golingon/terraproviders/google/5.25.0"
	"github.com/verifa/website/infra"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(fmt.Errorf("run: %w", err))
	}
}

type serviceStack struct {
	terra.Stack
	infra.Service
	Provider *google.Provider
	Backend  *infra.GCSBackend
}

func run(ctx context.Context) error {
	var (
		dryRun  bool
		destroy bool
	)

	flag.BoolVar(&dryRun, "dry-run", true, "dry run")
	flag.BoolVar(&destroy, "destroy", false, "destroy")
	flag.Parse()

	wf := sylt.NewWorkflow(
		sylt.WithWorkflowDryRun(dryRun),
		sylt.WithWorkflowDestroy(destroy),
	)
	defer func() {
		if err := wf.Cleanup(ctx); err != nil {
			slog.Error("running workflow cleanup", "error", err)
		}
	}()

	prNumber, ok := os.LookupEnv("GITHUB_PR")
	if !ok {
		return fmt.Errorf("GITHUB_PR not set")
	}

	prName := fmt.Sprintf("pr-%s", prNumber)

	svcPreview := sylt.Terra(prName+"-website", &serviceStack{
		Service: infra.NewService(infra.ServiceOpts{
			Name:     prName,
			MinScale: 0,
			MaxScale: 1,
		}),
		Backend: &infra.GCSBackend{
			Bucket: "verifa-website-tfstate",
			Prefix: prName,
		},
		Provider: &google.Provider{
			Project: terra.String("verifa-website"),
			Region:  terra.String("europe-north1"),
		},
	})
	if err := wf.Run(ctx, svcPreview); err != nil {
		return err
	}
	var stateErr error
	backendSvc := sylt.RequireResourceState(
		svcPreview.Stack.BackendService,
		&stateErr,
	)
	if stateErr != nil {
		return fmt.Errorf("preview environment state: %w", stateErr)
	}

	if err := wf.Run(ctx, &infra.URLMapAction{
		ID:      prName,
		Project: "verifa-website",
		URLMap:  "lb-prod-urlmap",
		Service: backendSvc.Id,
		Host:    fmt.Sprintf("%s.verifa.io", prName),
	}); err != nil {
		return fmt.Errorf("urlmap action: %w", err)
	}

	return nil
}
