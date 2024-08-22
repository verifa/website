package website

import (
	"bytes"
	"testing"

	"github.com/verifa/website/testutil"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func TestAdmonitionGolden(t *testing.T) {
	t.Parallel()
	type test struct {
		name   string
		source func() []byte
	}

	tests := []test{
		{
			name: "complex",
			source: func() []byte {
				return []byte(`# Title

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
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			md := goldmark.New(
				goldmark.WithParserOptions(
					parser.WithAutoHeadingID(),
				),
				goldmark.WithExtensions(
					&anchorExt{},
				),
			)
			context := parser.NewContext()
			buf := bytes.Buffer{}
			if err := md.Convert(test.source(), &buf, parser.WithContext(context)); err != nil {
				t.Fatalf("converting: %s", err)
			}
			testutil.Golden(t, buf.Bytes(), ".html")
		})
	}
}
