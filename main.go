package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	openMenuItem := fyne.NewMenuItem("Open...", func() {
		fmt.Println("Open...")
	})

	saveMenuItem := fyne.NewMenuItem("Save", func() {
		fmt.Println("Save...")
	})
	c.SaveMenuItem = saveMenuItem
	c.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("Save as...", c.saveAs(w))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	mainMenu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(mainMenu)
}

func (c *config) saveAs(win fyne.Window) func() {
	fmt.Println("saveAs")

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

			// Save the file
			write.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " " + write.URI().Name())

			c.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.Show()
	}
}
