package main

import (
	"bufio"
	"fmt"
	"os"
)

type bookmark struct {
	url     string
	title   string
	comment string
}

func addBookmark(fileLocation *string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Add a new bookmark (y/N)? ")
	response, err := reader.ReadString('\n')
	check(err)

	if response != "y\n" {
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

	newBookmark := bookmark{url: url, title: bookmarkTitle, comment: bookmarkComment}
	writeBookmark(newBookmark, fileLocation)

}

func writeBookmark(newBookmark bookmark, fileLocation *string) {
	file, err := os.OpenFile(*fileLocation, os.O_APPEND|os.O_WRONLY, 0666)
	check(err)

	_, err = file.WriteString(newBookmark.title[:len(newBookmark.title)-1] + ": " + newBookmark.url + "\t" + "// " + newBookmark.comment)
	check(err)
}
