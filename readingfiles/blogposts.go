package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err:= fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}
	
	var posts []Post
	for _, f := range dir {
		post, err := getPost(filesystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, fileName string) (Post, error) {
	postFile, err := filesystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readTrimmedLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	readBody := func() string {
		var bodyBuilder strings.Builder
		scanner.Scan() // ignore first line of ---
		for scanner.Scan() {
			bodyBuilder.WriteString(scanner.Text() + "\n")
		}
		return strings.TrimSuffix(bodyBuilder.String(), "\n")
	}

	return Post{Title: readTrimmedLine("Title: "), 
				Description: readTrimmedLine("Description: "),
				Tags: strings.Split(readTrimmedLine("Tags: "), ", "),
				Body: readBody()}, 
				nil
}