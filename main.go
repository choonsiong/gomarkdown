package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	CurrentFile   fyne.URI
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// Create a Fyne app
	fa := app.New()

	// Create a window for the app
	win := fa.NewWindow("Markdown Editor")

	// Get the user interfaces
	edit, preview := cfg.makeUI()

	// Set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// Show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (c *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	c.EditWidget = edit
	c.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (c *config) createMenuItems(w fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", func() {
		fmt.Println("Open...")
	})

	saveMenuItem := fyne.NewMenuItem("Save", func() {
		fmt.Println("Save...")
	})

	saveAsMenuItem := fyne.NewMenuItem("Save as...", func() {
		fmt.Println("Save as...")
	})

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	mainMenu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(mainMenu)
}
