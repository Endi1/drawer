package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"strconv"
	"strings"
)

func listBookmarks(filename *string) {
	contentString := getFileContent(filename)
	bookmarksToPrint := parseBookmarksFile(contentString)
	printBookmarks(bookmarksToPrint)
}

func getFileContent(filename *string) string {
	content, err := ioutil.ReadFile(*filename)
	check(err)
	contentString := string(content)
	return contentString
}

func parseBookmarksFile(contentString string) []bookmark {
	splitBookmarks := strings.Split(contentString, "\n\n")
	var parsedBookmarks []bookmark

	for _, element := range splitBookmarks[0 : len(splitBookmarks)-1] {
		secondaryPart := strings.Split(element, "\n")     // Splice each bookmark between title+url and the rest
		mainPart := strings.Split(secondaryPart[0], ": ") // Splice title+url in title and url
		mainPartTitleID := strings.Split(mainPart[0], ". ")

		bookmarkID, err := strconv.Atoi(mainPartTitleID[0])
		check(err)

		formattedBookmark := bookmark{id: bookmarkID, url: mainPart[1], title: mainPartTitleID[1], comment: secondaryPart[1]}
		parsedBookmarks = append(parsedBookmarks, formattedBookmark)
	}

	return parsedBookmarks

}

func printBookmarks(bookmarksToPrint []bookmark) {
	titleColor := color.New(color.FgCyan).SprintFunc()
	urlColor := color.New(color.FgBlue).SprintFunc()
	commentColor := color.New(color.FgWhite).SprintFunc()
	IDColor := color.New(color.FgYellow).SprintFunc()

	for _, element := range bookmarksToPrint {
		fmt.Printf("%s. %s: %s \n%4s\n\n", IDColor(element.id), titleColor(element.title), urlColor(element.url), commentColor(element.comment))
	}
}
