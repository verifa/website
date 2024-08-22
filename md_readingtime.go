package website

import (
	"bytes"
	"log/slog"
	"math"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var readingTimeKey = parser.NewContextKey()

// wordsPerMinute is the average number of words a person can read in a minute.
// It is used to calculate the reading time of a post.
const wordsPerMinute = 212

var _ goldmark.Extender = (*readingTimeExt)(nil)

type readingTimeExt struct{}

func (e *readingTimeExt) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&readingTimeTransformer{}, 0),
	))
}

var _ parser.ASTTransformer = (*readingTimeTransformer)(nil)

// readingTimeTransformer is a Goldmark AST transformer that calculates the
// reading time of a post.
type readingTimeTransformer struct{}

func (*readingTimeTransformer) Transform(
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

	minutes := math.Ceil(float64(totalWords) / wordsPerMinute)
	duration := time.Duration(minutes * float64(time.Minute))
	pc.Set(readingTimeKey, duration)
}
