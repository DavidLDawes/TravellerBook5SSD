package main

import "fyne.io/fyne/widget"

var (
	saveButton = widget.NewButton("Save", saveMe)
	loadButton = widget.NewButton("Load", loadMe)
	// We have plenty of setttings so break 'em into 2
	settings       *widget.Form
	secondSettings *widget.Form
	// and all the details in one panel after that
	details    *widget.Box
	noneString = "None"
)

func saveMe() {
}

func loadMe() {
}
