package main

import (
	"github.com/endi1/drawer/drawer"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"os/exec"
	"strings"
)

type tui struct {
	app          *tview.Application
	form         *tview.Form
	pages        *tview.Pages
	fileLocation *string
}

func Init(bookmarks *[]drawer.BookmarkObject, fileLocation *string) {
	app := tview.NewApplication()
	pages := tview.NewPages()
	t := tui{app: app, pages: pages, fileLocation: fileLocation}
	app.SetInputCapture(t.inputActions)
	listView := t.list(bookmarks)
	addView := t.addView()

	pages.AddPage("Bookmarks", listView, true, true)
	pages.AddPage("New Bookmark", addView, true, false)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func (t *tui) list(bookmarks *[]drawer.BookmarkObject) *tview.List {
	list := tview.NewList()

	for _, bookmark := range *bookmarks {
		list.AddItem(bookmark.GetTitle(), bookmark.GetURL(), 0, func() {
			command := exec.Command("xdg-open", bookmark.GetURL())
			err := command.Run()
			check(err)
		})
	}

	return list
}

func (t *tui) inputActions(e *tcell.EventKey) *tcell.EventKey {
	switch pressed_key := e.Rune(); pressed_key {
	case 'q':
		t.app.Stop()
	case 'a':
		t.pages.SwitchToPage("New Bookmark")
	}
	return e
}

func (t *tui) addView() *tview.Form {
	form := tview.NewForm()
	t.form = form
	form.AddInputField("URL", "", 20, nil, nil)
	form.AddInputField("Title", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)
	form.AddInputField("Tags (comma-separated)", "", 20, nil, nil)
	form.AddButton("Save", t.handleAddAction)
	form.AddButton("Quit", func() {
		t.app.Stop()
	})
	return form
}

func (t *tui) handleAddAction() {
	urlField := t.form.GetFormItemByLabel("URL")
	titleField := t.form.GetFormItemByLabel("Title")
	tagsField := t.form.GetFormItemByLabel("Tags (comma-separated)")
	commentField := t.form.GetFormItemByLabel("Comment")

	url := urlField.(*tview.InputField).GetText()
	title := titleField.(*tview.InputField).GetText()
	tagsText := tagsField.(*tview.InputField).GetText()
	comment := commentField.(*tview.InputField).GetText()

	tags := parseTagsString(tagsText)

	newBookmark := bookmark{id: 0, url: url, title: title, comment: comment, tags: tags}
	writeBookmark(newBookmark, t.fileLocation)
}

func parseTagsString(tagsString string) []string {
	splitTags := strings.Split(tagsString, ",")
	return splitTags
}
