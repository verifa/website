//go:build tools
// +build tools

package website

import (
	_ "github.com/a-h/templ/cmd/templ"
	_ "github.com/cosmtrek/air"
	_ "github.com/google/ko"

	// CI tools
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)
