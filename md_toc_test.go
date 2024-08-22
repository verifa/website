package website

import (
	"bytes"
	"testing"

	"github.com/verifa/website/testutil"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func TestTableOfContents(t *testing.T) {
	t.Parallel()
	md := goldmark.New(
		goldmark.WithParserOptions(
			// TOC requres headings to have IDs.
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(

			&tableOfContentsExt{},
		),
	)

	type test struct {
		name   string
		source func() []byte
		exp    tableOfContents
	}

	tests := []test{
		{
			name:   "short",
			source: func() []byte { return []byte("## Subtitle\n### Subsubtitle") },
			exp: tableOfContents{
				{
					level: 2,
					text:  "Subtitle",
					url:   "#subtitle",
					children: []*heading{
						{
							level: 3,
							text:  "Subsubtitle",
							url:   "#subsubtitle",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := parser.NewContext()
			buf := bytes.Buffer{}
			if err := md.Convert(tt.source(), &buf, parser.WithContext(context)); err != nil {
				t.Fatalf("converting: %s", err)
			}
			toc := tableOfContentsFromContext(context)
			testutil.AssertNoDiff(t, tt.exp, toc)
		})
	}
}
