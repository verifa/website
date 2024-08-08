package watcher

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

type WatchOptions struct {
	// RunOnStart indicates whether to run the function on start.
	RunOnStart bool
	// Name is the name of this watcher.
	// It is used for logging/debugging purposes.
	Name string
	// Dir is the directory to watch files from.
	Dir string
	// IncludeFiles is a list of files relative to Dir to watch.
	IncludeFiles []string
	// IncludeSuffix is a list of file suffixes to watch.
	// Format: ".ext", "_templ.go", etc.
	IncludeSuffix []string
	// ExcludeSuffix is a list of file suffixes to exclude.
	// Format: ".ext", "_templ.go", etc.
	ExcludeSuffix []string

	// Fn is the function to call when a file is modified.
	Fn func([]string)
	// Batch is the time to wait before sending a batch of changed files.
	// This avoids sending multiple events if a tool (like Templ) modifies
	// multiple files at once. It is better to batch them together.
	Batch time.Duration
}

func WatchFilesystem(ctx context.Context, opt WatchOptions) error {
	logger := slog.With("watcher", opt.Name)
	if opt.Dir == "" {
		opt.Dir = "."
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("creating fsnotify watcher: %w", err)
	}

	// Start istening for events.
	go func() {
		defer watcher.Close()

		batching := false
		var batch []string
		var batchTimer <-chan time.Time
		for {
			select {
			case <-ctx.Done():
				return
			case <-batchTimer:
				// Send buffer.
				batching = false
				localBatch := make([]string, len(batch))
				copy(localBatch, batch)
				go opt.Fn(localBatch)
				batch = nil
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
					if !shouldIncludeFile(opt, event.Name) {
						continue
					}
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
				logger.Error("watcher error", "error", err)
			}
		}
	}()

	if err := filepath.WalkDir(opt.Dir, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if dir.IsDir() {
			if shouldSkipDir(dir.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if !shouldIncludeFile(opt, path) {
			return nil
		}
		logger.Info("watching", "path", path)
		// Adding the same directory to the watcher multiple times is a no-op.
		// Hence, we can safely add the same directory multiple times.
		watchDir := filepath.Dir(path)
		if err := watcher.Add(watchDir); err != nil {
			return fmt.Errorf("watching %s: %w", watchDir, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	if opt.RunOnStart {
		go opt.Fn([]string{})
	}

	return nil
}

func shouldIncludeFile(opts WatchOptions, name string) bool {
	if len(opts.IncludeFiles) > 0 && slices.Contains(opts.IncludeFiles, name) {
		return true
	}
	for _, suffix := range opts.ExcludeSuffix {
		if strings.HasSuffix(name, suffix) {
			return false
		}
	}
	for _, suffix := range opts.IncludeSuffix {
		if strings.HasSuffix(name, suffix) {
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
	// These directories are ignored by the Go tool.
	if strings.HasPrefix(dir, ".") || strings.HasPrefix(dir, "_") {
		return true
	}
	return false
}
