package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
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

func addBookmark(fileLocation *string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Add a new bookmark (y/N)? ")
	response, err := reader.ReadString('\n')
	check(err)

	if response != "y\n" {
		return
	}

	fmt.Print("Add a link: ")
	bookmark, err := reader.ReadString('\n')
	check(err)

	fmt.Print("Add a title for that link: ")
	bookmarkTitle, err := reader.ReadString('\n')
	check(err)

	writeBookmark(bookmark, bookmarkTitle, fileLocation)
}

func writeBookmark(bookmark string, bookmarkTitle string, fileLocation *string) {
	file, err := os.OpenFile(*fileLocation, os.O_APPEND|os.O_WRONLY, 0666)
	check(err)

	_, err = file.WriteString(bookmarkTitle[:len(bookmarkTitle)-1] + ": " + bookmark)
	check(err)
}

func listBookmarks(filename *string) {
	content, err := ioutil.ReadFile(*filename)
	check(err)
	contentString := string(content)
	fmt.Print(contentString)
}
