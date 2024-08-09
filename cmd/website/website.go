package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"

	"github.com/verifa/website"
)

var (
	buildGitCommit = "dev"
	isProduction   bool
)

func main() {
	flag.BoolVar(
		&isProduction,
		"prod",
		false,
		"run the website in production mode",
	)
	flag.Parse()

	if err := run(); err != nil {
		slog.Error("running website", "error", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := website.Run(ctx, website.Site{
		Commit:       buildGitCommit,
		IsProduction: isProduction,
	}); err != nil {
		return err
	}
	return nil
}
