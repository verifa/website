package website

import (
	"net/url"
	"strings"
	"testing"

	"github.com/verifa/website/testutil"
)

func mustParseURL(t *testing.T, s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	return u
}

func TestBibtex(t *testing.T) {
	type test struct {
		name  string
		input func() string
		exp   references
	}
	tests := []test{
		{
			name: "article",
			input: func() string {
				return `@article{article,
author = {Forsgren, Nicole and Kalliamvakou, Eirini and Noda, Abi and Greiler, Michaela and Houck, Brian and Storey, Margaret-Anne},
year = {2024},
month = {01},
pages = {47-77},
title = {DevEx in Action: A study of its tangible impacts},
volume = {21},
journal = {Queue},
doi = {10.1145/3639443}
url = {https://dl.acm.org/doi/10.1145/3639443}
}`
			},
			exp: references{
				{
					Key:   "article",
					Index: 1,
					Type:  "article",
					Title: "DevEx in Action: A study of its tangible impacts",
					Authors: []author{
						{
							First: "Nicole",
							Last:  "Forsgren",
						},
						{
							First: "Eirini",
							Last:  "Kalliamvakou",
						},
						{
							First: "Abi",
							Last:  "Noda",
						},
						{
							First: "Michaela",
							Last:  "Greiler",
						},
						{
							First: "Brian",
							Last:  "Houck",
						},
						{
							First: "Margaret-Anne",
							Last:  "Storey",
						},
					},
					Year: "2024",
					URL: mustParseURL(
						t,
						"https://dl.acm.org/doi/10.1145/3639443",
					),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := readBibtex(strings.NewReader(tt.input()))
			if err != nil {
				t.Fatal(err)
			}
			testutil.AssertNoDiff(t, tt.exp, entries)
		})
	}
}
