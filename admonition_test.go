package website

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

func TestAdmonition(t *testing.T) {
	t.Parallel()
	baseRenderOpts := []renderer.Option{
		html.WithHardWraps(),
		html.WithXHTML(),
		// Render raw HTML coming from the post markdown.
		html.WithUnsafe(),
	}

	admonitionRenderer := goldmark.DefaultRenderer()
	admonitionRenderer.AddOptions(baseRenderOpts...)
	admonitionRenderer.AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&AdmonitionBodyRenderer{}, 0),
		),
	)
	fullRenderOpts := append(
		baseRenderOpts,
		renderer.WithNodeRenderers(
			util.Prioritized(&AdmonitionRenderer{
				Renderer: admonitionRenderer,
			}, 0),
		),
	)
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&AdmonitionTransformer{}, 0),
			),
		),
		goldmark.WithRendererOptions(fullRenderOpts...),
	)

	source := []byte(`

# Title

> [!NOTE]
> Highlights information that users should take into account, even when skimming.
> Another line.
>
> A break in the line.
>
> Then some ` + "`code blocks`" + `.

> [!WARNING]
> A warning this time.

`)
	buf := bytes.Buffer{}
	if err := md.Convert(source, &buf); err != nil {
		t.Fatalf("converting: %s", err)
	}
	t.Log(buf.String())
}
