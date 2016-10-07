package main

import (
	"strconv"
	"strings"
)

func parseBookmarksFile(filename *string) []bookmark {
	contentString := getFileContent(filename)
	splitBookmarks := strings.Split(contentString, "\n\n")
	var parsedBookmarks []bookmark

	for _, element := range splitBookmarks[0 : len(splitBookmarks)-1] {
		secondaryPart := strings.Split(element, "\n") // Splice each bookmark between title+url and the rest
		comment := parseBookmarkComment(secondaryPart[1])

		mainPart := strings.Split(secondaryPart[0], ": h") // Splice title+url in title and url
		mainPartTitleID := strings.Split(mainPart[0], ". ")

		bookmarkID, err := strconv.Atoi(mainPartTitleID[0])
		check(err)

		formattedBookmark := bookmark{id: bookmarkID, url: "h" + mainPart[1], title: mainPartTitleID[1], comment: comment}
		parsedBookmarks = append(parsedBookmarks, formattedBookmark)
	}

	return parsedBookmarks

}

func parseBookmarkComment(content string) string {

	if len(content) == 4 {
		return ""
	}
	return content[3:]
}
