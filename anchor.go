package website

import (
	"fmt"
	"log/slog"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var _ parser.ASTTransformer = (*ReadingTimeTransformer)(nil)

type AnchorTransformer struct{}

// Transform implements parser.ASTTransformer.
func (*AnchorTransformer) Transform(
	node *ast.Document,
	reader text.Reader,
	pc parser.Context,
) {
	if err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if n.Kind() == ast.KindHeading {
			heading := n.(*ast.Heading)
			rawID, ok := heading.AttributeString("id")
			if !ok {
				return ast.WalkStop, fmt.Errorf("no id found for heading %s",
					heading.Text(reader.Source()))
			}
			id, ok := rawID.([]byte)
			if !ok {
				return ast.WalkStop, fmt.Errorf("id is not a string: %v", rawID)
			}

			str := ast.NewString([]byte("#"))
			anchor := ast.NewLink()
			anchor.Destination = append([]byte("#"), id...)
			// IMPORTANT: Any classes used here need to be safelisted in
			// the tailwind.config.cjs file.
			anchor.SetAttributeString("class", []byte("ml-2 text-v-gray no-underline hover:text-v-lilac"))
			anchor.AppendChild(anchor, str)
			heading.AppendChild(heading, anchor)
		}
		return ast.WalkContinue, nil
	}); err != nil {
		slog.Error("adding anchor to header", "error", err)
		return
	}
}
