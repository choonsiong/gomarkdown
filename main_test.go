package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_makeUI(t *testing.T) {
	var testCfg config
	edit, preview := testCfg.makeUI()

	test.Type(edit, "testing")

	if preview.String() != "testing" {
		t.Error("Failed to update preview with edit content")
	}
}

func Test_Runapp(t *testing.T) {
	var testCfg config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test Markdown")

	edit, preview := testCfg.makeUI()
	testCfg.createMenuItems(testWin)
	testWin.SetContent(container.NewHSplit(edit, preview))
	testApp.Run()

	test.Type(edit, "testing")

	if preview.String() != "testing" {
		t.Error("Failed to run app")
	}
}
