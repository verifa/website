package website

import "sort"

// TagStatus is used for rendering the tag status to a user.
type TagStatus int

const (
	// TagStatusNone indicates the tag is not present in the filtered posts.
	// The tag should be disabled and not checked.
	TagStatusNone TagStatus = iota
	// TagStatusChecked indicates the tag is checked.
	// The tag should be clickable and checked.
	TagStatusChecked
	// TagStatusRemaining indicates the tag is present in the filtered posts.
	// The tag should be clickable and not checked.
	TagStatusRemaining
)

func sortTags(tags map[string]TagStatus) []string {
	sorted := make([]string, 0, len(tags))
	for tag := range tags {
		sorted = append(sorted, tag)
	}
	sort.Strings(sorted)
	return sorted
}

// func authorNames(authors []Member) []string {
// 	names := make([]string, 0, len(authors))
// 	for _, author := range authors {
// 		names = append(names, author.Name)
// 	}
// 	return names
// }

// FilterBlogPosts filters the blog posts by the given tags.
// It returns the filtered posts and the status of each tag.
func FilterBlogPosts(
	posts *Posts,
	filterTags []string,
) ([]*Post, map[string]TagStatus) {
	filtered := make([]*Post, 0)
	remainingTags := make(map[string]TagStatus, len(posts.Tags))
	defaultTagStatus := TagStatusNone
	if len(filterTags) == 0 {
		defaultTagStatus = TagStatusRemaining
	}
	for tag := range posts.Tags {
		remainingTags[tag] = defaultTagStatus
	}
	if len(filterTags) == 0 {
		return posts.Blog, remainingTags
	}

	// Iterate over each blog post.
	// Check all the filter tags are present in the post, else skip it.
	for _, post := range posts.Blog {
		filterPost := false
		for _, tag := range filterTags {
			if _, ok := post.Tags[tag]; !ok {
				filterPost = true
			}
		}
		if filterPost {
			continue
		}
		filtered = append(filtered, post)
		for tag := range post.Tags {
			remainingTags[tag] = TagStatusRemaining
		}
	}
	// Mark all tags we are filtering by as checked (because they are).
	for _, tag := range filterTags {
		remainingTags[tag] = TagStatusChecked
	}
	return filtered, remainingTags
}
