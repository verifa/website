package website

import (
	"bytes"
	"testing"

	"github.com/verifa/website/testutil"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func TestCitations(t *testing.T) {
	t.Parallel()

	refs := references{
		{
			Key:   "abc",
			Index: 1,
			URL:   mustParseURL(t, "http://localhost:666"),
		},
	}
	source := []byte(`# Title

This is a \cite{abc} to a reference.
`)

	md := goldmark.New(
		goldmark.WithExtensions(&citationExt{
			References: refs,
		}),
	)

	context := parser.NewContext()
	buf := bytes.Buffer{}
	if err := md.Convert(source, &buf, parser.WithContext(context)); err != nil {
		t.Fatalf("converting: %s", err)
	}
	testutil.Golden(t, buf.Bytes(), ".html")
}
