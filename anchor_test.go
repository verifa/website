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

func TestAnchor(t *testing.T) {
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
		renderer.WithNodeRenderers(),
	)
	fullRenderOpts := baseRenderOpts
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithASTTransformers(
				util.Prioritized(&AnchorTransformer{}, 0),
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
	context := parser.NewContext()
	buf := bytes.Buffer{}
	if err := md.Convert(source, &buf, parser.WithContext(context)); err != nil {
		t.Fatalf("converting: %s", err)
	}
	t.Log(buf.String())
}
