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

> [!NOTE]
> A totally heinoius note.

` + "```" + `bash
terraform init
` + "```" + `

> [!NOTE]
> If you are working with local modules then there is no need to run ` + "`terraform init`" + ` before the scan as all files are already present, but remote modules must be fetched in to the ` + "`.terraform`" + ` folder before a scan.

Simplest way to run a Trivy misconfiguration scan is to point it at your current folder:

` + "```" + `bash
trivy config .
` + "```" + `

`)
	buf := bytes.Buffer{}
	if err := md.Convert(source, &buf); err != nil {
		t.Fatalf("converting: %s", err)
	}
	t.Log(buf.String())
}
