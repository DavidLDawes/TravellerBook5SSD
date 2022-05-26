package main

import "fyne.io/fyne/widget"

type summaryStruct struct {
	tons  float32
	cost  float32
	cargo float32
}

var (
	summary      summaryStruct = summaryStruct{}
	summaryLabel               = *widget.NewLabel("Summary goes here")
)

func (s summaryStruct) init(form *widget.Form, box *widget.Box) {
	summary.cargo = getTons()
	summaryLabel.SetText("Hull tonnage: 200, unarmored")
	summaryLabel.Show()
}
