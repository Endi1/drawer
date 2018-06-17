package drawer

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type T struct {
	Bookmarks []YAMLBookmark
}

type YAMLBookmark struct {
	Title   string
	URL     string
	Comment string
	Tags    []string
}

func getfileContent(filename *string) *string {
	content, err := ioutil.ReadFile(*filename)
	if err != nil {
		// Handle error
	}
	contentString := string(content)
	return &contentString
}

func getUnmarshaledContent(fileLocation *string) *T {
	fileContent := getfileContent(fileLocation)
	t := T{}
	err := yaml.Unmarshal([]byte(*fileContent), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &t
}

func getMarshaledContent(unmarshaledContent *T) *[]byte {
	d, err := yaml.Marshal(unmarshaledContent)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &d
}

func GetWrittenBookmarks(fileLocation *string) *[]YAMLBookmark {
	t := getUnmarshaledContent(fileLocation)
	return &t.Bookmarks
}

func WriteBookmarkToFile(bookmark BookmarkObject, fileLocation *string) {
	unmarshaledContent := getUnmarshaledContent(fileLocation)
	yamlBookmark := YAMLBookmark{
		Title:   *bookmark.getTitle(),
		URL:     *bookmark.getURL(),
		Comment: *bookmark.getComment(),
		Tags:    *bookmark.getTags(),
	}
	unmarshaledContent.Bookmarks = append(unmarshaledContent.Bookmarks,
		yamlBookmark)

	d := getMarshaledContent(unmarshaledContent)
	stringToWrite := string(*d)

	file, err := os.Create(*fileLocation)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(stringToWrite)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
