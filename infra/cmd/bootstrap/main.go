package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/golingon/lingon/pkg/terra"
	"github.com/golingon/lingon/pkg/x/sylt"
	google "github.com/golingon/terraproviders/google/5.25.0"
	"github.com/verifa/website/infra"
)

type stack struct {
	terra.Stack
	infra.GitHubOIDC
	Provider *google.Provider
	Backend  *infra.GCSBackend
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.Error("running", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	wf := sylt.NewWorkflow()
	defer func() {
		if err := wf.Cleanup(ctx); err != nil {
			slog.Error("running workflow cleanup", "error", err)
		}
	}()

	ghOIDC := sylt.Terra("github-oidc", &stack{
		GitHubOIDC: infra.NewGitHubOIDC(infra.GitHubOIDCOpts{
			Project: "verifa-website",
		}),
		Backend: &infra.GCSBackend{
			Bucket: "verifa-website-tfstate",
			Prefix: "github-oidc",
		},
		Provider: &google.Provider{
			Project: terra.String("verifa-website"),
			Region:  terra.String("europe-north1"),
			Zone:    terra.String("europe-north1-a"),
		},
	})
	if err := wf.Run(ctx, ghOIDC); err != nil {
		return fmt.Errorf("running github-oidc: %w", err)
	}
	return nil
}
