package main

import (
	"log"

	"gopkg.in/yaml.v2"
)

type T struct {
	Bookmarks []Bookmark
}

type Bookmark struct {
	Title   string
	URL     string
	Comment string
	Tags    []string
}

func parseBookmarksFile(filename *string) []bookmark {
	t := T{}
	contentString := getFileContent(filename)
	err := yaml.Unmarshal([]byte(contentString), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var parsedBookmarks []bookmark

	for index, element := range t.Bookmarks {
		parsedBookmark := bookmark{
			id:      index + 1,
			title:   element.Title,
			url:     element.URL,
			comment: element.Comment,
			tags:    element.Tags,
		}
		parsedBookmarks = append(parsedBookmarks, parsedBookmark)
	}

	return parsedBookmarks
}
