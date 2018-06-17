package main

import (
	"fmt"
	"github.com/endi1/drawer/drawer"
	"github.com/fatih/color"
	"io/ioutil"
)

func listBookmarks(fileLocation *string) {
	bookmarksToPrint := drawer.GetWrittenBookmarkObjects(fileLocation)
	Init(bookmarksToPrint, fileLocation)
}

func getFileContent(filename *string) string {
	content, err := ioutil.ReadFile(*filename)
	check(err)
	contentString := string(content)
	return contentString
}

func printBookmarks(bookmarksToPrint []bookmark) {

	for _, element := range bookmarksToPrint {
		fmt.Printf(printBookmark(element) + "\n")
	}
}

func printBookmark(bookmarkToPrint bookmark) string {
	var stringToPrint string

	titleColor := color.New(color.FgCyan).SprintFunc()
	commentColor := color.New(color.FgWhite).SprintFunc()
	IDColor := color.New(color.FgYellow).SprintFunc()
	tagsColor := color.New(color.FgRed).SprintFunc()

	stringToPrint = IDColor(bookmarkToPrint.id) + ". " + titleColor(bookmarkToPrint.title) + ": " + bookmarkToPrint.url + "\n"

	if bookmarkToPrint.comment != "" {
		stringToPrint += commentColor("// "+bookmarkToPrint.comment) + "\n"
	}

	if len(bookmarkToPrint.tags) != 0 {
		stringToPrint += tagsColor(tagsToString(bookmarkToPrint.tags)) + "\n"
	}

	return stringToPrint
}
