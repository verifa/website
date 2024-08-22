package website

import (
	"bytes"
	"testing"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func TestReadingTime(t *testing.T) {
	t.Parallel()
	md := goldmark.New(
		goldmark.WithExtensions(&readingTimeExt{}),
	)

	type test struct {
		name    string
		source  func() []byte
		minutes float64
	}

	tests := []test{
		{
			name:    "short",
			source:  func() []byte { return []byte("This is a short post.") },
			minutes: 1,
		},
		{
			name: "long",
			source: func() []byte {
				words := make([][]byte, 10*wordsPerMinute)
				for i := range words {
					words[i] = []byte("word")
				}
				return bytes.Join(words, []byte(" "))
			},
			minutes: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := parser.NewContext()
			buf := bytes.Buffer{}
			if err := md.Convert(tt.source(), &buf, parser.WithContext(context)); err != nil {
				t.Fatalf("converting: %s", err)
			}
			readingTime := context.Get(readingTimeKey).(time.Duration)
			if readingTime.Minutes() != tt.minutes {
				t.Errorf(
					"reading time: got %f, want %f",
					readingTime.Minutes(),
					tt.minutes,
				)
			}
		})
	}
}
