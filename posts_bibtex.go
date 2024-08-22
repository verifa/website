package website

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/jschaf/bibtex"
	"github.com/jschaf/bibtex/ast"
)

type references []reference

func (r references) cite(key string) (reference, bool) {
	for _, ref := range r {
		if ref.Key == key {
			return ref, true
		}
	}
	return reference{}, false
}

type reference struct {
	// Key is the bibtex (or reference) key.
	Key string
	// Index is the index of the reference in the list of references.
	Index   int
	Type    string
	Title   string
	Authors []author
	Year    string
	URL     *url.URL
}

type author struct {
	First string
	Last  string
}

func referencesFromBibtex(postsFS embed.FS, path string) (references, error) {
	bibPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".bib"
	bibFile, err := postsFS.Open(bibPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("opening bib file %s: %w", bibPath, err)
	}
	return readBibtex(bibFile)
}

func readBibtex(in io.Reader) (references, error) {
	bib := bibtex.New(
		bibtex.WithResolvers(
			// NewAuthorResolver creates a resolver for the "author" field that
			// parses author names into an ast.Authors node.
			bibtex.NewAuthorResolver("author"),
			// RenderParsedTextResolver replaces ast.ParsedText with a
			// simplified rendering of ast.Text.
			bibtex.NewRenderParsedTextResolver(),
		),
	)
	bibFile, err := bib.Parse(in)
	if err != nil {
		return nil, err
	}
	entries, err := bib.Resolve(bibFile)
	if err != nil {
		return nil, err
	}
	references := make(references, len(entries))
	for i, entry := range entries {
		ref := reference{
			// References start at 1, not 0. Hence the +1.
			Index: i + 1,
			Key:   entry.Key,
			// TODO: validate the type
			Type: entry.Type,
		}
		title, err := entryTagAsString(entry, bibtex.FieldTitle)
		if err != nil {
			return nil, err
		}
		ref.Title = title

		authorsTag, ok := entry.Tags[bibtex.FieldAuthor]
		if !ok {
			return nil, fmt.Errorf(
				"missing field %q in entry %q",
				bibtex.FieldAuthor,
				entry.Key,
			)
		}
		authors, ok := authorsTag.(ast.Authors)
		if !ok {
			return nil, fmt.Errorf(
				"field %q is not text in entry %q",
				bibtex.FieldAuthor,
				entry.Key,
			)
		}
		ref.Authors = make([]author, len(authors))
		for i, member := range authors {
			ref.Authors[i] = author{
				First: member.First.(*ast.Text).Value,
				Last:  member.Last.(*ast.Text).Value,
			}
		}

		rawURL, err := entryTagAsString(entry, bibtex.Field("url"))
		if err != nil {
			return nil, err
		}

		year, err := entryTagAsString(entry, bibtex.FieldYear)
		if err != nil {
			return nil, err
		}
		ref.Year = year

		url, err := url.Parse(rawURL)
		if err != nil {
			return nil, fmt.Errorf("parsing URL %q: %w", rawURL, err)
		}
		ref.URL = url

		references[i] = ref
	}
	return references, nil
}

func entryTagAsString(entry bibtex.Entry, field bibtex.Field) (string, error) {
	tag, ok := entry.Tags[field]
	if !ok {
		return "", fmt.Errorf("missing field %q in entry %q", field, entry.Key)
	}
	text, ok := tag.(*ast.Text)
	if !ok {
		return "", fmt.Errorf(
			"field %q is not text in entry %q",
			field,
			entry.Key,
		)
	}
	return text.Value, nil
}
