package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/verifa/website"
)

var buildGitCommit = "dev"

func main() {
	var isProduction bool
	flag.BoolVar(
		&isProduction,
		"prod",
		false,
		"run the website in production mode",
	)
	flag.Parse()

	if err := website.Run(website.Site{
		Commit:       buildGitCommit,
		IsProduction: isProduction,
	}); err != nil {
		slog.Error("starting website", "error", err)
		os.Exit(1)
	}
}
