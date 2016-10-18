package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func tagsToString(tags []string) string {
	var tagsString string

	for i, element := range tags {
		if i != len(tags)-1 {
			tagsString = tagsString + "#" + element + " "
		} else {
			tagsString = tagsString + "#" + element
		}
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
		strconv.Itoa(newBookmark.id) + ". " + newBookmark.title + ": " + newBookmark.url + "\n" + "// " + newBookmark.comment + "\n" + tagsToString(newBookmark.tags) + "\n\n")
	check(err)
}
