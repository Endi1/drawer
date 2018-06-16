package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
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
	t := T{}
	contentString := getFileContent(fileLocation)
	err := yaml.Unmarshal([]byte(contentString), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	yamlBookmark := Bookmark{
		Title:   newBookmark.title,
		URL:     newBookmark.url,
		Tags:    newBookmark.tags,
		Comment: newBookmark.comment,
	}

	t.Bookmarks = append(t.Bookmarks, yamlBookmark)
	d, err := yaml.Marshal(&t)
	check(err)

	newString := string(d)

	file, err := os.Create(*fileLocation)
	check(err)

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(newString)
	check(err)
}
