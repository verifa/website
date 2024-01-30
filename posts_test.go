package website

import (
	"testing"
)

func TestPosts(t *testing.T) {
	t.Parallel()
	posts, err := ParsePosts(postsFS)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(posts)
}
