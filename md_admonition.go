package website

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var _ goldmark.Extender = (*admonitionExt)(nil)

// admonitionExt is a Goldmark extension for admonitions.
type admonitionExt struct{}

func (e *admonitionExt) Extend(m goldmark.Markdown) {
	rendr := renderer.NewRenderer(
		html.WithHardWraps(),
		// Allow raw HTML coming from the post markdown.
		html.WithUnsafe(),
		renderer.WithNodeRenderers(
			util.Prioritized(html.NewRenderer(), 1000),
			util.Prioritized(&emptyNodeRenderer{
				nodeKinds: []ast.NodeKind{admonitionNodeKind},
			}, 0),
		),
	)
	m.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&admonitionTransformer{}, 0),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&admonitionRenderer{
			Renderer: rendr,
		}, 0),
	))
}

type admonitionLevel string

const (
	AdmonitionTypeNote      admonitionLevel = "NOTE"
	AdmonitionTypeTip       admonitionLevel = "TIP"
	AdmonitionTypeImportant admonitionLevel = "IMPORTANT"
	AdmonitionTypeWarning   admonitionLevel = "WARNING"
	AdmonitionTypeCaution   admonitionLevel = "CAUTION"
)

func (a admonitionLevel) Title() (string, error) {
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

var _ parser.ASTTransformer = (*admonitionTransformer)(nil)

// admonitionTransformer transforms blockquotes into admonitions.
type admonitionTransformer struct{}

func (*admonitionTransformer) Transform(
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
		adLevel := []byte{}
		remainder := []ast.Node{}
		for child := p1.FirstChild(); child != nil; child = child.NextSibling() {
			text := child.Text(reader.Source())
			switch index {
			case 0:
				if child.Kind() != ast.KindText {
					return ast.WalkContinue, nil
				}
				if !bytes.Equal(text, []byte("[")) {
					return ast.WalkContinue, nil
				}
			case 1:
				if child.Kind() != ast.KindText {
					return ast.WalkContinue, nil
				}
				if text[0] != '!' {
					return ast.WalkContinue, nil
				}
				adLevel = text[1:]
			case 2:
				if child.Kind() != ast.KindText {
					return ast.WalkContinue, nil
				}
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

		admonition := &admonitionBlock{
			AdType: admonitionLevel(adLevel),
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

// admonitionNodeKind is a NodeKind for admonition nodes.
var admonitionNodeKind = ast.NewNodeKind("Admonition")

var _ ast.Node = (*admonitionBlock)(nil)

// admonitionBlock is a block node for an admonition.
type admonitionBlock struct {
	ast.BaseBlock

	Title  string
	AdType admonitionLevel
}

// Dump implements ast.Node.
func (a *admonitionBlock) Dump(source []byte, level int) {
	ast.DumpHelper(a, source, level, nil, nil)
}

// Kind implements ast.Node.
func (*admonitionBlock) Kind() ast.NodeKind {
	return admonitionNodeKind
}

var _ renderer.NodeRenderer = (*admonitionRenderer)(nil)

// admonitionRenderer renders an Admonition node to HTML.
type admonitionRenderer struct {
	Renderer renderer.Renderer
}

func (r *admonitionRenderer) RegisterFuncs(
	reg renderer.NodeRendererFuncRegisterer,
) {
	reg.Register(admonitionNodeKind, r.renderAside)
}

func (r *admonitionRenderer) renderAside(
	writer util.BufWriter,
	source []byte,
	n ast.Node,
	entering bool,
) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	aside := n.(*admonitionBlock)

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
	if err := render(context.TODO(), writer); err != nil {
		return ast.WalkStop, fmt.Errorf("rendering aside: %w", err)
	}
	return ast.WalkSkipChildren, nil
}

var _ renderer.NodeRenderer = (*emptyNodeRenderer)(nil)

// emptyNodeRenderer renders the body of nodes to nothing.
// It is used by a rendered so that certain nodes are allowed in the AST, but
// they do not produce any HTML.
// This is because the bodies of such nodes are given to something like templ
// which produces the final HTML.
//
// It is a bit weird, but the best/easiest way we came up with to combine Templ
// and Goldmark.
type emptyNodeRenderer struct {
	nodeKinds []ast.NodeKind
}

func (r *emptyNodeRenderer) RegisterFuncs(
	reg renderer.NodeRendererFuncRegisterer,
) {
	for _, kind := range r.nodeKinds {
		reg.Register(kind, r.renderNothing)
	}
}

func (*emptyNodeRenderer) renderNothing(
	writer util.BufWriter,
	source []byte,
	n ast.Node,
	entering bool,
) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}
