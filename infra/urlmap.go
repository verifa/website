package infra

import (
	"context"
	"fmt"
	"log/slog"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/golingon/lingon/pkg/x/sylt"
)

var _ sylt.Actioner = (*URLMapAction)(nil)

const URLMapActionType sylt.ActionType = "urlmap"

type URLMapAction struct {
	// Unique ID of this URLMap.
	ID string
	// Google cloud project.
	Project string
	// URLMap to update.
	URLMap string
	// Backend service to map to.
	Service string
	// Host (domain) to map to the service.
	Host string
}

func (a *URLMapAction) ActionName() string {
	return a.ID
}

func (a *URLMapAction) ActionType() sylt.ActionType {
	return URLMapActionType
}

func (a *URLMapAction) Run(ctx context.Context, opts sylt.RunOpts) error {
	log := slog.With(
		"action_name",
		a.ActionName(),
		"action_type",
		URLMapActionType,
		"action_phase",
		"run",
		"opts",
		opts,
	)

	return a.run(ctx, log, opts.Destroy, opts.DryRun)
}

func (a *URLMapAction) Cleanup(
	ctx context.Context,
	opts sylt.RunOpts,
) error {
	log := slog.With(
		"action_name",
		a.ActionName(),
		"action_type",
		URLMapActionType,
		"action_phase",
		"cleanup",
		"opts",
		opts,
	)
	if opts.DryRun {
		log.Info("skipping cleanup")
		return nil
	}
	return a.run(ctx, log, opts.Destroy, false)
}

func (a *URLMapAction) run(
	ctx context.Context,
	log *slog.Logger,
	destroy bool,
	dryRun bool,
) error {
	client, err := compute.NewUrlMapsRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("creating google cloud urlmap client: %w", err)
	}
	urlMap, err := client.Get(ctx, &computepb.GetUrlMapRequest{
		Project: a.Project,
		UrlMap:  a.URLMap,
	})
	if err != nil {
		return fmt.Errorf("getting url map: %w", err)
	}
	rule := urlMapRule{
		ID:      a.ID,
		Host:    a.Host,
		Service: a.Service,
	}
	// Alwaysu remove the URLMap rule before adding it back (if not destroy).
	// This ensures it gets updated if we made changes.
	removeURLMapRule(urlMap, rule.ID)
	if !destroy {
		addURLMapRule(urlMap, rule)
	}
	log.Info("updating url map", "rule", rule, "host_rules", urlMap.HostRules)

	if dryRun {
		log.Info("skipping update due to dry-run")
		return nil
	}

	op, err := client.Update(ctx, &computepb.UpdateUrlMapRequest{
		Project:        "verifa-website",
		UrlMap:         *urlMap.Name,
		UrlMapResource: urlMap,
	})
	if err != nil {
		return fmt.Errorf("updating url map: %w", err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for operation: %w", err)
	}
	return nil
}

type urlMapRule struct {
	ID      string
	Host    string
	Service string
}

func addURLMapRule(urlMap *computepb.UrlMap, rule urlMapRule) {
	urlMap.HostRules = append(urlMap.HostRules, &computepb.HostRule{
		Hosts:       []string{rule.Host},
		PathMatcher: p(rule.ID),
	})
	urlMap.PathMatchers = append(urlMap.PathMatchers, &computepb.PathMatcher{
		Name:           p(rule.ID),
		DefaultService: p(rule.Service),
	})
}

func removeURLMapRule(urlMap *computepb.UrlMap, ruleID string) {
	var newHostRules []*computepb.HostRule
	var newPathMatchers []*computepb.PathMatcher
	for _, hostRule := range urlMap.HostRules {
		if *hostRule.PathMatcher != ruleID {
			newHostRules = append(newHostRules, hostRule)
		}
	}
	for _, pathMatcher := range urlMap.PathMatchers {
		if *pathMatcher.Name != ruleID {
			newPathMatchers = append(newPathMatchers, pathMatcher)
		}
	}
	urlMap.HostRules = newHostRules
	urlMap.PathMatchers = newPathMatchers
}

func p[T any](t T) *T {
	return &t
}
