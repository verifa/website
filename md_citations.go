package website

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var citationRegExp = regexp.MustCompile(`\\cite{([A-Za-z0-9-]+)}`)

const citationRefIndex = 1

var _ goldmark.Extender = (*citationExt)(nil)

type citationExt struct {
	References references
}

func (e *citationExt) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(&citationParser{}, 0),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&citationRenderer{
			References: e.References,
		}, 500),
	))
}

var citationNodeKind = ast.NewNodeKind("Citation")

var _ ast.Node = (*citationNode)(nil)

// citationNode is a Goldmark AST node for a citation.
type citationNode struct {
	ast.BaseInline

	Reference string
}

func (n *citationNode) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

func (n *citationNode) Kind() ast.NodeKind {
	return citationNodeKind
}

var _ parser.InlineParser = (*citationParser)(nil)

// citationParser is a Goldmark inline parser for citations.
// It finds citations in the source markdown and replaces them with citation
// nodes.
type citationParser struct{}

func (p *citationParser) Trigger() []byte {
	return []byte{'\\'}
}

func (p *citationParser) Parse(
	parent ast.Node,
	block text.Reader,
	pc parser.Context,
) ast.Node {
	line, _ := block.PeekLine()
	matches := citationRegExp.FindSubmatch(line)
	if len(matches) == 0 {
		return nil
	}
	citeFull := matches[0]
	block.Advance(len(citeFull))

	citeRef := matches[citationRefIndex]
	return &citationNode{
		Reference: string(citeRef),
	}
	// for next := block.Peek(); next != '}'; {
	// 	fmt.Printf("CITATION: %s\n", string(next))
	// 	block.Advance(1)
	// 	next = block.Peek()
	// }

	// block.Advance(1)
	// fmt.Printf("CITATION: %s\n", string(block.Peek()))
	// block.Advance(1)
	// return &NodeCitation{}
}

var _ renderer.NodeRenderer = (*citationRenderer)(nil)

// citationRenderer renders citation nodes to HTML.
type citationRenderer struct {
	References references
}

func (r *citationRenderer) RegisterFuncs(
	reg renderer.NodeRendererFuncRegisterer,
) {
	reg.Register(citationNodeKind, r.renderCitation)
}

func (r *citationRenderer) renderCitation(
	w util.BufWriter,
	source []byte,
	n ast.Node,
	entering bool,
) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	citeNode, ok := n.(*citationNode)
	if !ok {
		return ast.WalkStop, errors.New("not a citation node")
	}
	ref, ok := r.References.cite(citeNode.Reference)
	if !ok {
		return ast.WalkStop, fmt.Errorf(
			"reference not found for %q",
			citeNode.Reference,
		)
	}

	if err := citationTemplate(ref).Render(context.TODO(), w); err != nil {
		return ast.WalkStop, fmt.Errorf("rendering citation: %w", err)
	}
	return ast.WalkContinue, nil
}
