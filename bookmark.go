package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type bookmark struct {
	id      int
	url     string
	title   string
	comment string
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

func removeFromFile(bookmarks []bookmark, fileLocation *string) {
	err := ioutil.WriteFile(*fileLocation, []byte(""), 0644)
	check(err)

	for _, element := range bookmarks {
		writeBookmark(element, fileLocation)
	}

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

	contentString := getFileContent(fileLocation)

	if contentString != "" {
		parsedBookmarks := parseBookmarksFile(fileLocation)
		bookmarkID = getLastID(parsedBookmarks) + 1
	} else {
		bookmarkID = 0
	}

	newBookmark := bookmark{id: bookmarkID, url: url, title: title, comment: comment}
	writeBookmark(newBookmark, fileLocation)

}

func writeBookmark(newBookmark bookmark, fileLocation *string) {
	file, err := os.OpenFile(*fileLocation, os.O_APPEND|os.O_WRONLY, 0666)
	check(err)

	_, err = file.WriteString(strconv.Itoa(newBookmark.id) + ". " + newBookmark.title + ": " + newBookmark.url + "\n\t" + "// " + newBookmark.comment + "\n\n")
	check(err)
}

func getLastID(parsedBookmarks []bookmark) int {
	lastID := parsedBookmarks[len(parsedBookmarks)-1].id
	return lastID
}
