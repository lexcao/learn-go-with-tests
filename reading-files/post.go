package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile fs.File) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tag string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tag)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buffer := &bytes.Buffer{}
	for scanner.Scan() {
		_, _ = fmt.Fprintln(buffer, scanner.Text())
	}
	body := strings.TrimSuffix(buffer.String(), "\n")
	return body
}
