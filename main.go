package main

import (
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"gopkg.in/src-d/go-git.v4"
)

func main() {
	// TODO Load Config File
	// TODO Git Operations
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/barakmich/tpom",
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}

	if err := uiMain(); err != nil {
		panic(err)
	}
}

func uiMain() error {
	// Search Box
	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetPlaceholder("Search...")
	input.SetBackgroundColor(tcell.ColorDefault)
	input.SetFieldBackgroundColor(tcell.ColorDefault)

	// List Box
	list := tview.NewList()
	list.AddItem("Foo", "", '\x00', nil)
	list.SetBorderPadding(1, 1, 5, 5)
	list.SetSelectedFocusOnly(true)

	// Layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(list, 0, 1, false)
	flex.AddItem(input, 3, 1, true)

	// Execute
	app := tview.NewApplication().SetRoot(flex, true).SetFocus(input)
	return app.Run()
}
