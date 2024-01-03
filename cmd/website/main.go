package main

import (
	"log/slog"
	"os"

	"github.com/verifa/website"
)

func main() {
	if err := website.Run(); err != nil {
		slog.Error("starting website", "error", err)
		os.Exit(1)
	}
}
