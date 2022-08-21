package main

import (
	"fyne.io/fyne/v2"
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

}
