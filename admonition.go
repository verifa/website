package website

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type AdmonitionType string

const (
	AdmonitionTypeNote      AdmonitionType = "NOTE"
	AdmonitionTypeTip       AdmonitionType = "TIP"
	AdmonitionTypeImportant AdmonitionType = "IMPORTANT"
	AdmonitionTypeWarning   AdmonitionType = "WARNING"
	AdmonitionTypeCaution   AdmonitionType = "CAUTION"
)

func (a AdmonitionType) Title() (string, error) {
	switch a {
	case AdmonitionTypeNote:
		return "Note", nil
	case AdmonitionTypeTip:
		return "Tip", nil
	case AdmonitionTypeImportant:
		return "Important", nil
	case AdmonitionTypeWarning:
		return "Warning", nil
	case AdmonitionTypeCaution:
		return "Caution", nil
	default:
		return "", fmt.Errorf("unknown admonition type: %s", a)
	}
}

var _ parser.ASTTransformer = (*AdmonitionTransformer)(nil)

type AdmonitionTransformer struct{}

// Transform implements parser.ASTTransformer.
func (*AdmonitionTransformer) Transform(
	node *ast.Document,
	reader text.Reader,
	pc parser.Context,
) {
	// You cannot remove a node that is being visited by the walker.
	// Hence, store the blockquotes which should be transformed into admonitions
	// so that once the walker has finished we can remove them.
	blockquotesToRemove := []*ast.Blockquote{}
	if err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		bq, ok := n.(*ast.Blockquote)
		if !ok {
			return ast.WalkContinue, nil
		}

		if !bq.HasChildren() {
			return ast.WalkContinue, nil
		}

		// Check first child is a paragraph.
		p1, ok := bq.FirstChild().(*ast.Paragraph)
		if !ok {
			return ast.WalkContinue, nil
		}
		index := 0
		isAdmonition := false
		admonitionType := []byte{}
		remainder := []ast.Node{}
		for child := p1.FirstChild(); child != nil; child = child.NextSibling() {
			if child.Kind() != ast.KindText {
				return ast.WalkContinue, nil
			}
			text := child.Text(reader.Source())
			switch index {
			case 0:
				if !bytes.Equal(text, []byte("[")) {
					return ast.WalkContinue, nil
				}
			case 1:
				if text[0] != '!' {
					return ast.WalkContinue, nil
				}
				admonitionType = text[1:]
			case 2:
				if !bytes.Equal(text, []byte("]")) {
					return ast.WalkContinue, nil
				}
				isAdmonition = true
			default:
				remainder = append(remainder, child)
			}
			index++
		}
		if !isAdmonition {
			return ast.WalkContinue, nil
		}

		admonition := &Admonition{
			AdType: AdmonitionType(admonitionType),
		}
		if len(remainder) > 0 {
			p := ast.NewParagraph()
			for _, child := range remainder {
				p.AppendChild(p, child)
			}
			admonition.AppendChild(admonition, p)
		}
		bq.RemoveChild(bq, p1)
		// Get remaining children of blockquote.
		// We cannot append the children directly to the aside because then the
		// NextSibling does not exist.
		// Instead, first add them to a slice and then append them to the aside
		// from the slice.
		children := make([]ast.Node, 0, bq.ChildCount())
		for child := bq.FirstChild(); child != nil; child = child.NextSibling() {
			children = append(children, child)
		}
		for _, child := range children {
			admonition.AppendChild(admonition, child)
		}
		// Insert the admonition before the blockquote and add blockquote to
		// list to be removed from the ast later.
		bq.Parent().InsertBefore(bq.Parent(), bq, admonition)
		blockquotesToRemove = append(blockquotesToRemove, bq)

		return ast.WalkSkipChildren, nil
	}); err != nil {
		slog.Error("transforming admonition", "error", err)
		return
	}

	// After tree has been walked, remove the blockquotes.
	for _, bq := range blockquotesToRemove {
		bq.Parent().RemoveChild(bq.Parent(), bq)
	}
}

// KindAside is a NodeKind of the Aside node.
var KindAside = ast.NewNodeKind("Aside")

var _ ast.Node = (*Admonition)(nil)

type Admonition struct {
	ast.BaseBlock

	Title  string
	AdType AdmonitionType
}

// Dump implements ast.Node.
func (a *Admonition) Dump(source []byte, level int) {
	ast.DumpHelper(a, source, level, nil, nil)
}

// Kind implements ast.Node.
func (*Admonition) Kind() ast.NodeKind {
	return KindAside
}

var _ renderer.NodeRenderer = (*AdmonitionRenderer)(nil)

type AdmonitionRenderer struct {
	Renderer renderer.Renderer
}

func (r *AdmonitionRenderer) RegisterFuncs(
	reg renderer.NodeRendererFuncRegisterer,
) {
	reg.Register(
		KindAside,
		func(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
			if !entering {
				return ast.WalkContinue, nil
			}
			aside := n.(*Admonition)

			title := aside.Title
			if title == "" {
				var err error
				title, err = aside.AdType.Title()
				if err != nil {
					return ast.WalkStop, fmt.Errorf(
						"getting admonition title: %w",
						err,
					)
				}
			}

			buf := bytes.Buffer{}
			if err := r.Renderer.Render(&buf, source, aside); err != nil {
				return ast.WalkStop, fmt.Errorf("rendering aside: %w", err)
			}
			var render templ.ComponentFunc
			switch aside.AdType {
			case AdmonitionTypeNote:
				render = admonitionNote(title, buf.Bytes()).Render
			case AdmonitionTypeTip:
				render = admonitionTip(title, buf.Bytes()).Render
			case AdmonitionTypeImportant:
				render = admonitionImportant(title, buf.Bytes()).Render
			case AdmonitionTypeWarning:
				render = admonitionWarning(title, buf.Bytes()).Render
			case AdmonitionTypeCaution:
				render = admonitionCaution(title, buf.Bytes()).Render
			default:
				return ast.WalkStop, fmt.Errorf(
					"unknown admonition type: %s",
					aside.AdType,
				)
			}
			if err := render(context.Background(), writer); err != nil {
				return ast.WalkStop, fmt.Errorf("rendering aside: %w", err)
			}
			return ast.WalkSkipChildren, nil
		},
	)
}

var _ renderer.NodeRenderer = (*AdmonitionBodyRenderer)(nil)

type AdmonitionBodyRenderer struct{}

func (r *AdmonitionBodyRenderer) RegisterFuncs(
	reg renderer.NodeRendererFuncRegisterer,
) {
	reg.Register(
		KindAside,
		func(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
			return ast.WalkContinue, nil
		},
	)
}
