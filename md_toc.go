package website

import (
	"fmt"
	"log/slog"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var _ goldmark.Extender = (*tableOfContentsExt)(nil)

type tableOfContentsExt struct{}

func (e *tableOfContentsExt) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&tableOfContentsTransformer{}, 0),
	))
}

var tableOfContentsKey = parser.NewContextKey()

func tableOfContentsFromContext(pc parser.Context) tableOfContents {
	toc, ok := pc.Get(tableOfContentsKey).(tableOfContents)
	if !ok {
		return nil
	}
	return toc
}

type tableOfContents []*heading

type heading struct {
	level    int
	text     string
	url      string
	children []*heading
}

var _ parser.ASTTransformer = (*tableOfContentsTransformer)(nil)

// tableOfContentsTransformer is a Goldmark AST transformer that stores the
// table of contents inside the goldmark parser context.
// It does not modify the AST.
type tableOfContentsTransformer struct{}

func (*tableOfContentsTransformer) Transform(
	node *ast.Document,
	reader text.Reader,
	pc parser.Context,
) {
	toc := tableOfContents{}
	var parent *heading
	if err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if n.Kind() == ast.KindHeading {
			h := n.(*ast.Heading)
			if h.Level == 1 {
				return ast.WalkStop, fmt.Errorf("post should not have a level 1 heading in its")
			}
			// Check that the first heading is level 2.
			if parent == nil && h.Level != 2 {
				return ast.WalkStop, fmt.Errorf("first header must always be level 2 (##)")
			}
			// Get the "id" of the heading to use as the URL.
			rawID, ok := h.AttributeString("id")
			if !ok {
				return ast.WalkStop, fmt.Errorf("no id found for heading %s",
					h.Text(reader.Source()))
			}
			id, ok := rawID.([]byte)
			if !ok {
				return ast.WalkStop, fmt.Errorf("id is not a string: %v", rawID)
			}
			head := heading{
				level: h.Level,
				text:  string(h.Text(reader.Source())),
				url:   fmt.Sprintf("#%s", string(id)),
			}
			if h.Level == 2 {
				toc = append(toc, &head)
				parent = &head
			} else {
				parent.children = append(parent.children, &head)
			}
		}
		return ast.WalkContinue, nil
	}); err != nil {
		slog.Error(
			"collecting table of contents",
			"error",
			err,
			"meta",
			node.Text(reader.Source()),
		)
		return
	}
	pc.Set(tableOfContentsKey, toc)
}
