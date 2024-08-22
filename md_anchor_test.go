package website

import (
	"bytes"
	"testing"

	"github.com/verifa/website/testutil"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func TestAnchorGolden(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		source func() []byte
	}

	tests := []test{
		{
			name: "basic",
			source: func() []byte {
				return []byte(`# Title

A paragraph.

## Subtitle

Another paragraph.

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
