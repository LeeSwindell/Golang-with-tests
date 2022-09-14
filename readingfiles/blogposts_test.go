package blogposts_test

import (
	blogposts "posts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
    )
	
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{
				Title: "Post 1", 
				Description: "Description 1",
				Tags: []string{"tdd", "go"},
				Body: `Hello
World`,
			}
	assertPosts(t, got, want)

	got = posts[1]
	want = blogposts.Post{
			Title: "Post 2",
			Description: "Description 2",
			Tags: []string{"rust", "borrow-checker"},
			Body: `B
L
M`,
	} 
	assertPosts(t, got, want)
}

func assertPosts(t testing.TB, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}