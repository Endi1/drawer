package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileLocation := flag.String("f", ".mypocket", "location of bookmark file")
	addBookmarkBool := flag.Bool("a", false, "add a new bookmark")

	flag.Parse()
	fmt.Println("file:" + *fileLocation)

	if *addBookmarkBool {
		addBookmark(fileLocation)
		return
	}
	createFileOrListBookmarks(fileLocation)

}

func createFileOrListBookmarks(fileLocation *string) {
	if _, err := os.Stat(*fileLocation); os.IsNotExist(err) {
		_, err := os.Create(*fileLocation)
		check(err)
		addBookmark(fileLocation)
	}
	listBookmarks(fileLocation)
}

func listBookmarks(filename *string) {
	content, err := ioutil.ReadFile(*filename)
	check(err)
	contentString := string(content)
	c := color.New(color.FgCyan)
	c.Println(contentString)
	parseBookmarksFile(contentString)
}

func parseBookmarksFile(contentString string) {
	bookmarksStrings := strings.Split(contentString, "\n\n")
	var bookmarksToPrint []bookmark

	for _, element := range bookmarksStrings {
		elementSpliced := strings.Split(element, "\n")
		spliceForTitle := strings.Split(elementSpliced[0], ": ")
		tempBookmark := bookmark{url: spliceForTitle[1], title: spliceForTitle[0], comment: elementSpliced[1]}
		bookmarksToPrint = append(bookmarksToPrint, tempBookmark)
	}
	fmt.Println(bookmarksToPrint)
}
