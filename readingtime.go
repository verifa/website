package website

import (
	"bytes"
	"log/slog"
	"math"
	"time"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var readingTimeKey = parser.NewContextKey()

var _ parser.ASTTransformer = (*ReadingTimeTransformer)(nil)

type ReadingTimeTransformer struct{}

// Transform implements parser.ASTTransformer.
func (*ReadingTimeTransformer) Transform(
	node *ast.Document,
	reader text.Reader,
	pc parser.Context,
) {
	totalWords := 0
	if err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if !n.HasChildren() {
			text := n.Text(reader.Source())
			totalWords += len(bytes.Fields(text))
		}
		return ast.WalkContinue, nil
	}); err != nil {
		slog.Error("calculating reading time", "error", err)
		return
	}

	// Calculate reading time using 212 words per minute, taken from Hugo.
	minutes := math.Ceil(float64(totalWords) / 212.0)
	duration := time.Duration(minutes * float64(time.Minute))
	pc.Set(readingTimeKey, duration)
}
