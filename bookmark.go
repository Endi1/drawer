package main

import (
	"bufio"
	"fmt"
	"os"
)

type bookmark struct {
	id      int
	url     string
	title   string
	comment string
	tags    []string
}

func deleteBookmark(bookmarkToDelete bookmark, bookmarks []bookmark, fileLocation *string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(printBookmark(bookmarkToDelete))
	fmt.Printf("Delete bookmark (y/N)? ")
	response, err := reader.ReadString('\n')
	check(err)

	if response == "y\n" {
		removeFromFile(bookmarks, fileLocation)
	}
	return
}

func addBookmark(fileLocation *string) {
	var bookmarkID int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Add a new bookmark (Y/n)? ")
	response, err := reader.ReadString('\n')
	check(err)

	if response == "n\n" {
		return
	}

	fmt.Print("Add a link: ")
	url, err := reader.ReadString('\n')
	check(err)
	url = url[:len(url)-1]

	fmt.Print("Add a title for that link: ")
	title, err := reader.ReadString('\n')
	check(err)
	title = title[:len(title)-1]

	fmt.Print("Add a comment: ")
	comment, err := reader.ReadString('\n')
	check(err)
	comment = comment[:len(comment)-1]

	var tags []string
	tags = getTags(reader, &tags)

	contentString := getFileContent(fileLocation)

	if contentString != "" {
		parsedBookmarks := parseBookmarksFile(fileLocation)
		bookmarkID = getLastID(parsedBookmarks) + 1
	} else {
		bookmarkID = 0
	}

	newBookmark := bookmark{id: bookmarkID, url: url, title: title, comment: comment, tags: tags}
	writeBookmark(newBookmark, fileLocation)

}

func getTags(reader *bufio.Reader, tags *[]string) []string {
	fmt.Print("Add a tag (empty to stop): ")
	tag, err := reader.ReadString('\n')
	check(err)
	if tag == "\n" {
		return *tags
	}
	tag = tag[:len(tag)-1]
	mytags := *tags
	*tags = append(mytags, tag)
	return getTags(reader, tags)
}

func getLastID(parsedBookmarks []bookmark) int {
	lastID := parsedBookmarks[len(parsedBookmarks)-1].id
	return lastID
}
