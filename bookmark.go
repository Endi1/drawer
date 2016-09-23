package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bookmark struct {
	id      int
	url     string
	title   string
	comment string
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

	fmt.Print("Add a title for that link: ")
	bookmarkTitle, err := reader.ReadString('\n')
	check(err)

	fmt.Print("Add a comment: ")
	bookmarkComment, err := reader.ReadString('\n')

	contentString := getFileContent(fileLocation)

	if contentString != "" {
		parsedBookmarks := parseBookmarksFile(contentString)
		bookmarkID = getLastID(parsedBookmarks) + 1
	} else {
		bookmarkID = 0
	}

	newBookmark := bookmark{id: bookmarkID, url: url, title: bookmarkTitle, comment: bookmarkComment}
	writeBookmark(newBookmark, fileLocation)

}

func writeBookmark(newBookmark bookmark, fileLocation *string) {
	file, err := os.OpenFile(*fileLocation, os.O_APPEND|os.O_WRONLY, 0666)
	check(err)

	_, err = file.WriteString(strconv.Itoa(newBookmark.id) + ". " + newBookmark.title[:len(newBookmark.title)-1] + ": " + newBookmark.url + "\t" + "// " + newBookmark.comment + "\n")
	check(err)
}

func getLastID(parsedBookmarks []bookmark) int {
	lastID := parsedBookmarks[len(parsedBookmarks)-1].id
	return lastID
}
