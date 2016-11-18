package main

import (
	"regexp"
	"strconv"
	"strings"
)

func parseBookmarksFile(filename *string) []bookmark {
	contentString := getFileContent(filename)
	re := regexp.MustCompile(`\d+(\S\s?)*`)
	splitBookmarks := re.FindAllString(contentString, -1)

	var parsedBookmarks []bookmark
	for _, element := range splitBookmarks {
		var comment string
		var tags []string

		secondaryPart := strings.Split(element, "\n") // Splice each bookmark between title+url and the rest

		if len(secondaryPart) == 2 {
			comment, tags = getCommentOrTags(secondaryPart[1])
		} else {
			comment = parseBookmarkComment(secondaryPart[1])
			tags = tagsToSplice(secondaryPart[2])
		}

		formattedBookmark := bookmark{id: parseID(secondaryPart[0]), url: parseURL(secondaryPart[0]), title: parseTitle(secondaryPart[0]), comment: comment, tags: tags}
		parsedBookmarks = append(parsedBookmarks, formattedBookmark)
	}

	return parsedBookmarks

}

func parseTitle(mainPart string) string {
	re := regexp.MustCompile(`\d[.](\s\S+)*[:]\s?h`)
	title := re.FindString(mainPart)

	re = regexp.MustCompile(`\d[.](\s?)`)
	title = re.ReplaceAllLiteralString(title, "")

	re = regexp.MustCompile(`\s?[:]\s?h`)
	return re.ReplaceAllLiteralString(title, "")
}

func parseURL(mainPart string) string {
	re := regexp.MustCompile(`:(\s?https?://\S+)`)
	url := re.FindString(mainPart)

	re = regexp.MustCompile(`:(\s?)h`)
	return re.ReplaceAllLiteralString(url, "h")
}

func parseID(mainPart string) int {
	re := regexp.MustCompile(`\d+[.]`)

	id := re.FindString(mainPart)
	id = strings.Replace(id, ".", "", -1)

	bookmarkID, err := strconv.Atoi(id)
	check(err)

	return bookmarkID
}

func getCommentOrTags(secondaryPart string) (string, []string) {
	comment, err := regexp.MatchString("^(//)", secondaryPart)
	check(err)

	if comment {
		return parseBookmarkComment(secondaryPart), []string{}
	}
	return "", tagsToSplice(secondaryPart)

}

func parseBookmarkComment(content string) string {
	re := regexp.MustCompile(`//(\s?\S+)*`)
	comment := re.FindString(content)

	return strings.Replace(comment, "//", "", -1)
}

func tagsToSplice(tagsString string) []string {
	re := regexp.MustCompile(`#([a-zA-Z]|\s)*`)
	tags := re.FindAllString(tagsString, -1)

	for i, tag := range tags {
		tags[i] = strings.Replace(tag, "#", "", -1)
	}

	return tags

}
