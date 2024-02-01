package main

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

type WatchOptions struct {
	Dir     string
	Include []string

	Fn    func([]string)
	Batch time.Duration
}

func WatchFilesystem(ctx context.Context, opt WatchOptions) error {
	if opt.Dir == "" {
		opt.Dir = "."
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("creating fsnotify watcher: %w", err)
	}

	// Start istening for events.
	go func() {
		batching := false
		batch := []string{}
		var batchTimer <-chan time.Time
		for {
			select {
			case <-ctx.Done():
				if err := watcher.Close(); err != nil {
					slog.Error("closing watcher", "error", err)
					return
				}
				slog.Info("watcher closed")
				return
			case <-batchTimer:
				// Send buffer.
				batching = false
				go opt.Fn(batch)
				batch = []string{}
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					batch = append(batch, event.Name)
					if batching {
						continue
					}
					batching = true
					batchTimer = time.After(opt.Batch)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				slog.Error("watcher error", "error", err)
			}
		}
	}()

	if err := filepath.WalkDir(opt.Dir, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if dir.IsDir() && shouldSkipDir(path) {
			return filepath.SkipDir
		}
		if !shouldIncludeFile(opt, path) {
			return nil
		}
		slog.Info("watching", "file", path)
		if err := watcher.Add(path); err != nil {
			return fmt.Errorf("watching %s: %w", path, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}

func shouldIncludeFile(opts WatchOptions, name string) bool {
	// include := []string{".templ", ".go", ".css", ".md"}
	for _, ext := range opts.Include {
		if strings.HasSuffix(name, ext) {
			return true
		}
	}
	return false
}

func shouldSkipDir(dir string) bool {
	if dir == "." {
		return false
	}
	if dir == "vendor" || dir == "node_modules" {
		return true
	}
	name := filepath.Dir(dir)
	// These directories are ignored by the Go tool.
	if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
		return true
	}
	return false
}
