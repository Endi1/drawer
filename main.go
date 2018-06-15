package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileLocation := flag.String("f", ".mydrawer", "location of bookmark file")
	addBookmarkBool := flag.Bool("a", false, "add a new bookmark")
	searchBookmarkID := flag.String("i", "", "get the bookmark with id")
	deleteBookmarkID := flag.String("d", "", "delete bookmark with this id")
	tagToSearch := flag.String("t", "", "get bookmarks with this tag")
	copyBookmark := flag.String("c", "", "[bookmark id] copy URL for bookmark with this id")

	flag.Parse()
	fmt.Println("file:" + *fileLocation)

	if *addBookmarkBool {
		addBookmark(fileLocation)
		return
	}

	if *searchBookmarkID != "" {
		searchBookmarkIDInt, err := strconv.Atoi(*searchBookmarkID)
		check(err)

		getBookmarkByID(searchBookmarkIDInt, fileLocation)
		return
	}

	if *deleteBookmarkID != "" {
		deleteBookmarkIDInt, err := strconv.Atoi(*deleteBookmarkID)
		check(err)

		deleteBookmarkByID(deleteBookmarkIDInt, fileLocation)
		return
	}

	if *tagToSearch != "" {
		searchBookmarksByTag(tagToSearch, fileLocation)
		return
	}

	if *copyBookmark != "" {
		copyBookmarkIDInt, err := strconv.Atoi(*copyBookmark)
		check(err)

		copyBookmarkByID(copyBookmarkIDInt, fileLocation)
		return
	}

	createFileOrListBookmarks(fileLocation)

}

func copyBookmarkByID(id int, fileLocation *string) {
	var bookmarkToCopy bookmark
	bookmarks := parseBookmarksFile(fileLocation)
	for _, element := range bookmarks {
		if element.id == id {
			bookmarkToCopy = element
		}
	}
	clipboard.WriteAll(bookmarkToCopy.url)

	fmt.Println("Copied bookmark:")
	fmt.Printf(printBookmark(bookmarkToCopy))

	return
}

func deleteBookmarkByID(id int, fileLocation *string) {
	var bookmarkToDelete bookmark
	bookmarks := parseBookmarksFile(fileLocation)
	for i, element := range bookmarks {
		if element.id == id {
			bookmarkToDelete = element
			bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
		}
	}
	deleteBookmark(bookmarkToDelete, bookmarks, fileLocation)
}

func getBookmarkByID(id int, fileLocation *string) {
	bookmarks := parseBookmarksFile(fileLocation)
	for _, element := range bookmarks {
		if element.id == id {
			fmt.Printf(printBookmark(element))
			return
		}
	}
}

func createFileOrListBookmarks(fileLocation *string) {
	if _, err := os.Stat(*fileLocation); os.IsNotExist(err) {
		_, err := os.Create(*fileLocation)
		check(err)
		addBookmark(fileLocation)
	}
	listBookmarks(fileLocation)
}

func searchBookmarksByTag(tag *string, fileLocation *string) {
	bookmarks := parseBookmarksFile(fileLocation)
	for _, bookmark := range bookmarks {
		for _, element := range bookmark.tags {
			if element == *tag {
				fmt.Printf(printBookmark(bookmark))
			}
		}
	}
}
