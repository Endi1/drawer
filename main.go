package main

import (
	"flag"
	"fmt"
	"os"
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
