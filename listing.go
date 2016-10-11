package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
)

func listBookmarks(filename *string) {
	bookmarksToPrint := parseBookmarksFile(filename)
	printBookmarks(bookmarksToPrint)
}

func getFileContent(filename *string) string {
	content, err := ioutil.ReadFile(*filename)
	check(err)
	contentString := string(content)
	return contentString
}

func printBookmarks(bookmarksToPrint []bookmark) {

	for _, element := range bookmarksToPrint {
		fmt.Printf(printBookmark(element))
	}
}

func printBookmark(bookmarkToPrint bookmark) string {
	var stringToPrint string

	titleColor := color.New(color.FgCyan).SprintFunc()
	urlColor := color.New(color.FgBlue).SprintFunc()
	commentColor := color.New(color.FgWhite).SprintFunc()
	IDColor := color.New(color.FgYellow).SprintFunc()
	tagsColor := color.New(color.FgRed).SprintFunc()

	stringToPrint = fmt.Sprintf("%s. %s: %s \n%4s\n%s\n\n", IDColor(bookmarkToPrint.id), titleColor(bookmarkToPrint.title), urlColor(bookmarkToPrint.url), commentColor(bookmarkToPrint.comment), tagsColor(tagsToString(bookmarkToPrint.tags)))

	return stringToPrint
}
