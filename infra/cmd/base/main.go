package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"

	"github.com/golingon/lingon/pkg/terra"
	"github.com/golingon/lingon/pkg/x/sylt"
	google "github.com/golingon/terraproviders/google/5.25.0"
	"github.com/verifa/website/infra"
)

type lbStack struct {
	terra.Stack
	infra.LoadBalancer
	Provider *google.Provider
	Backend  *infra.GCSBackend
}

type serviceStack struct {
	terra.Stack
	infra.Service
	Provider *google.Provider
	Backend  *infra.GCSBackend
}

type defaultBackendStack struct {
	terra.Stack
	infra.DefaultBackend
	Provider *google.Provider
	Backend  *infra.GCSBackend
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	var (
		apply   bool
		destroy bool
	)

	flag.BoolVar(&apply, "apply", false, "apply (not dry-run)")
	flag.BoolVar(&destroy, "destroy", false, "destroy")
	flag.Parse()

	wf := sylt.NewWorkflow(
		sylt.WithWorkflowDryRun(!apply),
		sylt.WithWorkflowDestroy(destroy),
	)
	defer func() {
		if err := wf.Cleanup(ctx); err != nil {
			slog.Error("running workflow cleanup", "error", err)
		}
	}()

	svcDefault := sylt.Terra("service-default", &defaultBackendStack{
		DefaultBackend: infra.NewDefaultBackend(),
		Backend: &infra.GCSBackend{
			Bucket: "verifa-website-tfstate",
			Prefix: "default-service",
		},
		Provider: &google.Provider{
			Project: terra.String("verifa-website"),
			Region:  terra.String("europe-north1"),
			Zone:    terra.String("europe-north1-a"),
		},
	})
	if err := wf.Run(ctx, svcDefault); err != nil {
		return err
	}
	defaultBackend, ok := svcDefault.Stack.DefaultBackend.Backend.State()
	if !ok {
		return fmt.Errorf("getting default backend state")
	}

	st := sylt.Terra("loadbalancer", &lbStack{
		LoadBalancer: infra.NewLoadBalancer(infra.LoadBalancerOpts{
			DefaultService: defaultBackend.Id,
		}),
		Backend: &infra.GCSBackend{
			Bucket: "verifa-website-tfstate",
			Prefix: "loadbalancer",
		},
		Provider: &google.Provider{
			Project: terra.String("verifa-website"),
			Region:  terra.String("europe-north1"),
			Zone:    terra.String("europe-north1-a"),
		},
	})
	if err := wf.Run(ctx, st); err != nil {
		return err
	}

	svcProd := sylt.Terra("service-prod", &serviceStack{
		Service: infra.NewService(infra.ServiceOpts{
			Name:     "prod-website",
			MinScale: 1,
			MaxScale: 10,
		}),
		Backend: &infra.GCSBackend{
			Bucket: "verifa-website-tfstate",
			Prefix: "prod-service",
		},
		Provider: &google.Provider{
			Project: terra.String("verifa-website"),
			Region:  terra.String("europe-north1"),
			Zone:    terra.String("europe-north1-a"),
		},
	})
	if err := wf.Run(ctx, svcProd); err != nil {
		return err
	}

	var stateErr error
	urlMap := sylt.RequireResourceState(
		st.Stack.URLMap,
		&stateErr,
	)
	prodService := sylt.RequireResourceState(
		svcProd.Stack.BackendService,
		&stateErr,
	)
	if stateErr != nil {
		return stateErr
	}

	if err := wf.Run(ctx, &infra.URLMapAction{
		ID:      "pathmatcher-0",
		Project: "verifa-website",
		URLMap:  urlMap.Name,
		Service: prodService.Id,
		Host:    "verifa.io",
	}); err != nil {
		return fmt.Errorf("urlmap action: %w", err)
	}

	return nil
}
