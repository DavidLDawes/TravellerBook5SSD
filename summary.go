package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

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
	summary.cargo = getCargo()
	summary.tons = float32(hullDetails.tons)
	summaryLabel.SetText(fmt.Sprintf("Hull tonnage: %d - %.1f tons used, leaving %.1f tons for cargo",
		hullDetails.tons, getTons(), getCargo()))
	summaryLabel.Show()
	box.Children = append(box.Children, &summaryLabel)
}

func (s summaryStruct) update() {
	summary.cargo = getCargo()
	summary.tons = float32(hullDetails.tons)
	summaryLabel.SetText(fmt.Sprintf("Hull tonnage: %d - %.1f tons used, leaving %.1f tons for cargo",
		hullDetails.tons, getTons(), getCargo()))
	summaryLabel.Show()
}
