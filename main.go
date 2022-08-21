package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io"
	"strings"
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
	cfg.createMenuItems(win)

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
	openMenuItem := fyne.NewMenuItem("Open...", c.openFunc(w))

	saveMenuItem := fyne.NewMenuItem("Save", func() {
		fmt.Println("Save...")
	})
	c.SaveMenuItem = saveMenuItem
	c.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("Save as...", c.saveAsFunc(w))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	mainMenu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(mainMenu)
}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (c *config) openFunc(win fyne.Window) func() {
	fmt.Println("openFunc")

	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil {
				return
			}
			defer read.Close()

			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			c.EditWidget.SetText(string(data))
			c.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " " + read.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (c *config) saveAsFunc(win fyne.Window) func() {
	fmt.Println("saveAsFunc")

	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if write == nil {
				// User cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with .md extension!", win)
				return
			}

			defer write.Close()

			// Save the file
			write.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = write.URI()
			win.SetTitle(win.Title() + " " + write.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}
