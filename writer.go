package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func tagsToString(tags []string) string {
	var tagsString string
	for _, element := range tags {
		tagsString = tagsString + "#" + element + " "
	}

	return tagsString
}

func removeFromFile(bookmarks []bookmark, fileLocation *string) {
	err := ioutil.WriteFile(*fileLocation, []byte(""), 0644)
	check(err)

	for _, element := range bookmarks {
		writeBookmark(element, fileLocation)
	}

}

func writeBookmark(newBookmark bookmark, fileLocation *string) {
	file, err := os.OpenFile(*fileLocation, os.O_APPEND|os.O_WRONLY, 0666)
	check(err)

	_, err = file.WriteString(
		strconv.Itoa(newBookmark.id) + ". " + newBookmark.title + ": " + newBookmark.url + "\n\t" + "// " + newBookmark.comment + "\n\t" + tagsToString(newBookmark.tags) + "\n\n")
	check(err)
}
