package website

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

//go:embed posts/*
var postsFS embed.FS

type PostType string

const (
	PostTypeBlog      PostType = "Blog"
	PostTypePodcast   PostType = "Podcast"
	PostTypeEvent     PostType = "Event"
	PostTypeJob       PostType = "Job"
	PostTypeCaseStudy PostType = "Case"
)

type Posts struct {
	All   []*Post
	Blog  []*Post
	Jobs  []*Post
	Cases []*Post

	// Index is a map of slugs to posts.
	Index map[string]*Post

	// ByAuthor is a map of author IDs to posts.
	ByAuthor map[string][]*Post

	// Tags is a map of Tags and posts that have that tag.
	Tags map[string][]*Post
}

func (p Posts) Featured() []*Post {
	totalFeatured := 3
	numFeatured := 0

	featured := make([]*Post, 0, totalFeatured)
	for _, post := range p.Blog {
		if numFeatured >= totalFeatured {
			break
		}
		// Ignore hidden or non-featured posts.
		if post.Hidden || !post.Featured {
			continue
		}
		featured = append(featured, post)
		numFeatured++
	}
	return featured
}

type Post struct {
	Slug         string
	Type         PostType
	Active       bool
	Title        string
	Subheading   string
	Authors      []Member
	Tags         map[string]struct{}
	Date         time.Time
	LastMod      time.Time
	ReadingTime  time.Duration
	PreviewImage string
	Image        string
	Featured     bool
	Hidden       bool
	Body         []byte

	// TableOfContents stores the headings in the post.
	TableOfContents     tableOfContents
	ShowTableOfContents bool

	// References stores the references defined in the post.
	References []reference

	SimilarPosts []*Post
}

func (p Post) URL() string {
	switch p.Type {
	case PostTypeBlog, PostTypeEvent, PostTypePodcast:
		return fmt.Sprintf("/blog/%s/", p.Slug)
	case PostTypeJob:
		return fmt.Sprintf("/careers/%s/", p.Slug)
	case PostTypeCaseStudy:
		return fmt.Sprintf("/work/%s/", p.Slug)
	default:
		return "/404/"
	}
}

func defaultPostsMarkdown(references references) goldmark.Markdown {
	baseRenderOpts := []renderer.Option{
		html.WithHardWraps(),
		// Allow raw HTML coming from the post markdown.
		html.WithUnsafe(),
	}

	md := goldmark.New(
		goldmark.WithParserOptions(
			// Anchor and Table Of Contents requres headings to have IDs.
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
			highlighting.NewHighlighting(
				highlighting.WithStyle("catppuccin-frappe"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			// Table of Contents should come before anchor to avoid anchor's "#"
			// being included in the header.
			&tableOfContentsExt{},
			&admonitionExt{},
			&anchorExt{},
			&citationExt{
				References: references,
			},
			&readingTimeExt{},
		),
		goldmark.WithRendererOptions(baseRenderOpts...),
	)
	return md
}

func ParsePosts(postsFS embed.FS) (*Posts, error) {
	posts := Posts{
		Index:    make(map[string]*Post),
		ByAuthor: make(map[string][]*Post),
		Tags:     make(map[string][]*Post),
	}
	if err := fs.WalkDir(
		postsFS,
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".md" {
				return nil
			}
			slug := strings.TrimSuffix(filepath.Base(path), ".md")

			// Add references to context to make it available to the
			references, err := referencesFromBibtex(postsFS, path)
			if err != nil {
				return fmt.Errorf("getting references from bibtex: %w", err)
			}

			md := defaultPostsMarkdown(references)
			context := parser.NewContext()
			contents, err := postsFS.ReadFile(path)
			if err != nil {
				return fmt.Errorf("reading file %s: %w", path, err)
			}
			var buf bytes.Buffer
			if err := md.Convert(contents, &buf, parser.WithContext(context)); err != nil {
				return fmt.Errorf("converting %s: %w", path, err)
			}
			metadata := meta.Get(context)
			if metadata == nil {
				return fmt.Errorf("no metadata found in %s", path)
			}
			readingTime, ok := context.Get(readingTimeKey).(time.Duration)
			if !ok {
				return fmt.Errorf("no reading time found in %s", path)
			}

			post, err := newPost(slug, metadata, readingTime, buf.Bytes())
			if err != nil {
				return fmt.Errorf("creating post from %s: %w", path, err)
			}
			if !post.Active {
				return nil
			}
			// Add references to post.
			post.References = references
			post.TableOfContents = tableOfContentsFromContext(context)

			posts.All = append(posts.All, post)
			posts.Index[post.Slug] = post
			return nil
		},
	); err != nil {
		return nil, fmt.Errorf("walking posts: %w", err)
	}

	// Sort all posts.
	sort.SliceStable(posts.All, func(i, j int) bool {
		return posts.All[i].Date.After(posts.All[j].Date)
	})
	// Organise posts by type and add post to tags.
	// We want to add tags after we have sorted the posts.
	for _, post := range posts.All {
		switch post.Type {
		case PostTypeBlog, PostTypeEvent, PostTypePodcast:
			posts.Blog = append(posts.Blog, post)
		case PostTypeJob:
			posts.Jobs = append(posts.Jobs, post)
			// Don't add job posts to tags because we don't want them to appear
			// as related posts.
			continue
		case PostTypeCaseStudy:
			posts.Cases = append(posts.Cases, post)
		default:
			return nil, fmt.Errorf("unknown post type %s", post.Type)
		}
		for tag := range post.Tags {
			posts.Tags[tag] = append(posts.Tags[tag], post)
		}
		// Add posts by author.
		for _, author := range post.Authors {
			posts.ByAuthor[author.ID] = append(posts.ByAuthor[author.ID], post)
		}
		// Add first three related posts.
		// This is not very efficient so if it affects startup time we could
		// look to improve it.
		for _, otherPost := range posts.All {
			if len(post.SimilarPosts) >= 3 {
				break
			}
			for tag := range post.Tags {
				if _, ok := otherPost.Tags[tag]; ok {
					post.SimilarPosts = append(post.SimilarPosts, otherPost)
					break
				}
			}
		}

	}
	// Old posts use tags we don't want to show.
	// Filter those out.
	filteredTags := []string{
		"rancher",
		"conference",
		"data-science",
		"jfrog",
		"machine-learning",
		"oss-compliance",
		"software-development",
	}
	for _, tag := range filteredTags {
		delete(posts.Tags, tag)
	}
	return &posts, nil
}

func newPost(
	slug string,
	metadata map[string]interface{},
	readingTime time.Duration,
	body []byte,
) (*Post, error) {
	if strings.Contains(slug, " ") {
		return nil, errors.New("slug contains space")
	}
	postType, ok := metadata["type"].(string)
	if !ok {
		return nil, errors.New("invalid post type")
	}
	active, ok := metadata["active"].(bool)
	if !ok {
		// Assume true.
		active = true
	}
	title, ok := metadata["title"].(string)
	if !ok {
		return nil, errors.New("invalid post title")
	}
	subheading, ok := metadata["subheading"].(string)
	if !ok {
		return nil, errors.New("invalid post subheading")
	}
	rawAuthors, ok := metadata["authors"].([]interface{})
	if !ok {
		return nil, errors.New("invalid post authors")
	}
	authorIDs := make([]string, len(rawAuthors))
	for i := range rawAuthors {
		authorIDs[i], ok = rawAuthors[i].(string)
		if !ok {
			return nil, fmt.Errorf(
				"invalid post author %d: %v",
				i,
				rawAuthors[i],
			)
		}
	}
	authors := make([]Member, len(authorIDs))
	for i, id := range authorIDs {
		member, ok := Crew[id]
		if !ok {
			return nil, fmt.Errorf("invalid post author %s", id)
		}
		authors[i] = member
	}
	rawTags, ok := metadata["tags"].([]interface{})
	if !ok {
		return nil, errors.New("invalid post tags")
	}
	tags := make(map[string]struct{}, len(rawTags))
	for i := range rawTags {
		strTag, ok := rawTags[i].(string)
		if !ok {
			return nil, fmt.Errorf(
				"invalid post tag %d: %v",
				i,
				rawTags[i],
			)
		}
		strTag = strings.ReplaceAll(strings.ToLower(strTag), " ", "-")
		tags[strTag] = struct{}{}
	}
	rawDate, ok := metadata["date"].(string)
	if !ok {
		return nil, errors.New("invalid post date")
	}
	date, err := time.Parse(time.DateOnly, rawDate)
	if err != nil {
		return nil, fmt.Errorf("parsing post date: %w", err)
	}
	// LastMod is optional.
	rawLastMod, ok := metadata["lastMod"].(string)
	var lastMod time.Time
	if ok {
		var err error
		lastMod, err = time.Parse(time.DateOnly, rawLastMod)
		if err != nil {
			return nil, fmt.Errorf("parsing post lastMod: %w", err)
		}
	}
	previewImage, ok := metadata["previewImage"].(string)
	if !ok {
		previewImage = ""
	}
	image, ok := metadata["image"].(string)
	if !ok {
		return nil, errors.New("invalid post image")
	}
	featured, ok := metadata["featured"].(bool)
	if !ok {
		// Assume false.
		featured = false
	}
	showTOC, ok := metadata["toc"].(bool)
	if !ok {
		// Assume false.
		showTOC = false
	}
	hidden, ok := metadata["hidden"].(bool)
	if !ok {
		// Assume false.
		hidden = false
	}

	return &Post{
		Slug:                slug,
		Type:                PostType(postType),
		Active:              active,
		Title:               title,
		Subheading:          subheading,
		Authors:             authors,
		Tags:                tags,
		Date:                date,
		LastMod:             lastMod,
		ReadingTime:         readingTime,
		PreviewImage:        previewImage,
		Image:               image,
		Featured:            featured,
		ShowTableOfContents: showTOC,
		Hidden:              hidden,
		Body:                body,
	}, nil
}
