package website

import (
	"testing"
)

func TestPosts(t *testing.T) {
	posts, err := ParsePosts(postsFS)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(posts)
}
