package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

const (
	tlA int = iota
	tlB
	tlC
	tlD
	tlE
	tlF
	tlG
	tlH
	tlJ
	tlK
	tlL
)

type tech struct {
	level  string
	offset int
}

var (
	defaultTech       = "A"
	defaultTechOffset = 0

	techDetails = tech{defaultTech, defaultTechOffset}
	techLevels  = []string{"A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L"}
	detailTech  = widget.NewLabel("Tech Level")
	techSelect  = widget.NewSelect(techLevels, nothing)
)

func (t tech) init(form *widget.Form, box *widget.Box) {
	techSelect.SetSelected(defaultTech)
	techSelect.PlaceHolder = defaultTech
	techSelect.Selected = defaultTech
	t.techChanged(defaultTech)
	techSelect.OnChanged = t.techChanged
	form.AppendItem(widget.NewFormItem("Tech Level", techSelect))

	box.Children = append(box.Children, detailTech)
}

func (t tech) techChanged(tlSelected string) {
	t.level = tlSelected
	t.offset = techToOffset(tlSelected)
	detailTech.SetText(fmt.Sprintf("Tech Level %s\n", t.level))
	detailTech.Refresh()
	detailTech.Show()
}

func (t tech) getTons() float32 {

	return 0.0
}

func offsetToTech(tlOffsetIn int) (result string) {
	result = "A"
	switch tlOffsetIn {
	default:
	case tlA:
		result = "A"
	case tlB:
		result = "B"
	case tlC:
		result = "C"
	case tlD:
		result = "D"
	case tlE:
		result = "E"
	case tlF:
		result = "F"
	case tlG:
		result = "G"
	case tlH:
		result = "H"
	case tlJ:
		result = "J"
	case tlK:
		result = "K"
	case tlL:
		result = "L"
	}

	return
}

func techToOffset(techIn string) (result int) {
	switch techIn {
	default:
	case "A":
		result = tlA
	case "B":
		result = tlB
	case "C":
		result = tlC
	case "D":
		result = tlD
	case "E":
		result = tlE
	case "F":
		result = tlF
	case "G":
		result = tlG
	case "H":
		result = tlH
	case "J":
		result = tlJ
	case "K":
		result = tlK
	case "L":
		result = tlL
	}

	return
}

func nothing(parm string) {
}