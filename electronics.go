package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type electronicsDetails struct {
	description string
	tlMin       int
	massByTL    []float32
	costByTL    []float32
	military    bool
	advanced    bool
	array       bool
	computer    int32
}

var (
	noElectronics = electronicsDetails{
		"None", 0,
		[]float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		false, false, false, 0,
	}
	LowCommElectronics = electronicsDetails{
		"Low End Comms", 0,
		[]float32{1, 1, .8, .7, .6, .5, .4, .3, .2, .1, .1, .1, .1, .1, .1},
		[]float32{2, 1.8, 1.5, 1.2, 1, .8, .7, .6, .5, .4, .3, .2, .1, .1, .1},
		false, false, false, 1,
	}
	LowSensorElectronics = electronicsDetails{
		"Low End Sensors", 0,
		[]float32{2, 2, 2, 1.9, 1.8, 1.7, 1.6, 1.5, 1.4, 1.3, 1.2, 1.1, 1, .9, .8},
		[]float32{3, 2.5, 2.4, 2.3, 2.2, 2.1, 2, 1.9, 1.8, 1.7, 1.6, 1.5, 1.4, 1.3, 1.2},
		false, false, false, 1,
	}
	LowElectronics = electronicsDetails{
		"Low End Comms & Sensors", 0,
		[]float32{3, 3, 2.8, 2.6, 2.5, 2.3, 2.1, 2, 1.8, 1.7, 1.5, 1.4, 1.2, 1.1, 1},
		[]float32{3, 3, 2.8, 2.6, 2.5, 2.3, 2.1, 2, 1.8, 1.7, 1.5, 1.4, 1.2, 1.1, 1},
		false, false, false, 1,
	}
	CommercialElectronics = electronicsDetails{
		"Commercial Comms & Sensors", 0,
		[]float32{6, 6, 5.7, 5.5, 5.2, 5, 4.6, 4.2, 3.9, 3.5, 3.3, 3, 2.7, 2.5, 2.2},
		[]float32{6, 6, 5.7, 5.5, 5.2, 5, 4.6, 4.2, 3.9, 3.5, 3.3, 3, 2.7, 2.5, 2.2},
		false, false, false, 2,
	}
	AdvancedElectronics = electronicsDetails{
		"Advanced Comms & Sensors", 3,
		[]float32{50, 50, 50, 12, 11, 10, 10, 10, 9.2, 8.6, 8, 7.6, 7.3, 7, 6.8},
		[]float32{50, 50, 50, 12, 11, 10, 10, 10, 9.2, 8.6, 8, 7.6, 7.3, 7, 6.8},
		false, true, false, 3,
	}
	CompactElectronics = electronicsDetails{
		"Compact Comms & Sensors", 5,
		[]float32{100, 100, 100, 100, 100, 3, 2, 2, 1.8, 1.7, 1.5, 1.4, 1.3, 1.2, 1.1},
		[]float32{100, 100, 100, 100, 100, 3, 2, 2, 1.8, 1.7, 1.5, 1.4, 1.3, 1.2, 1.1},
		false, false, false, 2,
	}
	AdvancedCompactElectronics = electronicsDetails{
		"Advanced Compact Comms & Sensors", 7,
		[]float32{200, 200, 200, 200, 200, 200, 200, 4, 4, 3.8, 3.5, 3.2, 3, 2.7, 2.5},
		[]float32{200, 200, 200, 200, 200, 200, 200, 4, 4, 3.8, 3.5, 3.2, 3, 2.7, 2.5},
		false, true, false, 3,
	}
	MilitaryElectronics = electronicsDetails{
		"Military Comms & Sensors", 1,
		[]float32{300, 30, 24, 20, 20, 19.5, 19.5, 19.2, 19, 18.5, 18, 17.5, 17, 16.5, 16},
		[]float32{300, 30, 24, 20, 20, 19.5, 19.5, 19.2, 19, 18.5, 18, 17.5, 17, 16.5, 16},
		false, false, false, 4,
	}
	AdvancedMilitaryElectronics = electronicsDetails{
		"Advanced Military Comms & Sensors", 4,
		[]float32{400, 400, 400, 400, 50, 48, 45, 40, 36, 32, 30, 28, 27, 26, 25},
		[]float32{400, 400, 400, 400, 50, 48, 45, 40, 36, 32, 30, 28, 27, 26, 25},
		false, false, false, 5,
	}
	CommercialWithArrayElectronics = electronicsDetails{
		"Commercial Comms & Sensors with Sensor Array", 0,
		[]float32{6, 6, 5.7, 5.5, 5.2, 5, 4.6, 4.2, 3.9, 3.5, 3.3, 3, 2.7, 2.5, 2.2},
		[]float32{6, 6, 5.7, 5.5, 5.2, 5, 4.6, 4.2, 3.9, 3.5, 3.3, 3, 2.7, 2.5, 2.2},
		false, false, true, 4,
	}
	AdvancedWithArrayElectronics = electronicsDetails{
		"Advanced Comms & Sensors with Sensor Array", 3,
		[]float32{50, 50, 50, 12, 11, 10, 10, 10, 9.2, 8.6, 8, 7.6, 7.3, 7, 6.8},
		[]float32{50, 50, 50, 12, 11, 10, 10, 10, 9.2, 8.6, 8, 7.6, 7.3, 7, 6.8},
		false, true, true, 5,
	}
	MilitaryWithArrayElectronics = electronicsDetails{
		"Military Comms & Sensors with Sensor Array", 1,
		[]float32{300, 30, 24, 20, 20, 19.5, 19.5, 19.2, 19, 18.5, 18, 17.5, 17, 16.5, 16},
		[]float32{300, 30, 24, 20, 20, 19.5, 19.5, 19.2, 19, 18.5, 18, 17.5, 17, 16.5, 16},
		false, false, true, 6,
	}
	AdvancedMilitaryWithArrayElectronics = electronicsDetails{
		"Advanced Military Comms & Sensors with Sensor Array", 4,
		[]float32{300, 30, 24, 20, 20, 19.5, 19.5, 19.2, 19, 18.5, 18, 17.5, 17, 16.5, 16},
		[]float32{300, 30, 24, 20, 20, 19.5, 19.5, 19.2, 19, 18.5, 18, 17.5, 17, 16.5, 16},
		false, false, true, 7,
	}

	generalSelections = []electronicsDetails{
		noElectronics, LowCommElectronics, LowSensorElectronics, LowElectronics,
		CommercialElectronics, AdvancedElectronics, CompactElectronics, AdvancedCompactElectronics,
		MilitaryElectronics, AdvancedMilitaryElectronics,
	}
	capitalSelections = []electronicsDetails{
		CommercialElectronics, CommercialWithArrayElectronics,
		AdvancedElectronics, AdvancedWithArrayElectronics,
		MilitaryElectronics, MilitaryWithArrayElectronics,
		AdvancedMilitaryElectronics, AdvancedMilitaryWithArrayElectronics,
	}
	electronics = CommercialElectronics

	computerSelections = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
	}

	computerMass = []float32{
		1.0, 2, 3.0, 4.0, 5.0, 7.0, 9.0, 11.0, 12.0, 13.0,
		15.0, 18.0, 20.0, 22.0, 24.0, 26.0, 28.0, 33.0, 36.0, 40.0, 45.0, 50.0, 56.0,
	}

	computerCost = []float32{
		1.0, 2, 12.0, 20.0, 30.0, 50.0, 70.0, 80.0, 90.0, 100.0,
		115.0, 125.0, 140.0, 150.0, 165.0, 180.0, 200.0, 215.0, 230.0, 250.0, 265.0, 280.0, 300.0,
	}

	electronicsSelect *widget.Select
	computerSelect    *widget.Select

	electronicsLabel = *widget.NewLabel("Electronics details go here")
	computerLabel    = *widget.NewLabel("Computer details go here")
)

func (e electronicsDetails) init(form *widget.Form, box *widget.Box) {
	box.Children = append(box.Children, &electronicsLabel, &computerLabel)

	electronicsSelect = widget.NewSelect(e.getSelections(generalSelections), stringValuedNothing)
	electronicsSelect.SetSelected(e.description)
	electronicsSelect.Selected = e.description
	electronicsSelect.Show()

	computerSelect = widget.NewSelect(computerSelections, stringValuedNothing)
	computerSelect.SetSelected("2")
	computerSelect.Selected = "2"
	computerSelect.Show()

	electronicsLabel.SetText(e.updateElectronics())
	electronicsLabel.Show()

	computerLabel.SetText(e.updateComputer())
	computerLabel.Show()

	electronicsSelect.OnChanged = e.electronicsChanged
	computerSelect.OnChanged = electronics.computerChanged

	form.AppendItem(widget.NewFormItem("Electronics Suite", electronicsSelect))
	form.AppendItem(widget.NewFormItem("Computer", computerSelect))
}

func (e electronicsDetails) updateElectronics() (electronicsDetails string) {
	electronicsDetails = fmt.Sprintf("E Suite: %s, %.1f tons costing %.1f",
		e.description,
		e.massByTL[techDetails.offset],
		e.costByTL[techDetails.offset],
	)

	return
}

func (e electronicsDetails) updateComputer() (computerDetails string) {
	computerDetails = fmt.Sprintf("Computer: %d, %.1f tons costing %.1f",
		electronics.computer,
		computerMass[electronics.computer],
		computerCost[electronics.computer],
	)

	return
}

func (e electronicsDetails) computerChanged(newSelection string) {
	for _, comp := range computerSelections {
		if comp == newSelection {
			convert, err := strconv.ParseInt(comp, 10, 16)
			if err == nil {
				electronics.computer = int32(convert)
				computerLabel.SetText(e.updateComputer())
				summary.update()
			}

			break
		}
	}
}

func (e electronicsDetails) electronicsChanged(newSelection string) {
	id := e.electronicsIndex(newSelection)
	if id > -1 {
		electronics.computer = e.getElectronicSelections()[id].computer
		e.massByTL = e.getElectronicSelections()[id].massByTL
		e.costByTL = e.getElectronicSelections()[id].costByTL
		e.description = newSelection
		e.advanced = e.getElectronicSelections()[id].advanced
		e.array = e.getElectronicSelections()[id].array
		e.military = e.getElectronicSelections()[id].military
		e.tlMin = e.getElectronicSelections()[id].tlMin

		electronicsLabel.SetText(e.updateElectronics())
		electronicsLabel.Show()
		if hullDetails.isCapital() {
			electronicsSelect = widget.NewSelect(e.getSelections(capitalSelections), stringValuedNothing)
		} else {
			electronicsSelect = widget.NewSelect(e.getSelections(generalSelections), stringValuedNothing)
		}
		summary.update()
	}
}

func (e electronicsDetails) getElectronicSelections() (available []electronicsDetails) {
	available = make([]electronicsDetails, 0)
	for _, nextDetails := range generalSelections {
		if nextDetails.tlMin <= techDetails.offset {
			available = append(available, nextDetails)
		}
	}

	return
}

func (e electronicsDetails) electronicsIndex(selection string) int {
	for index, match := range e.getElectronicSelections() {
		if match.description == selection {
			return index
		}
	}
	return -1
}

func (e electronicsDetails) getDetails(selection string) (details electronicsDetails, err error) {
	for _, match := range e.getElectronicSelections() {
		if match.description == selection {
			details = match

			return
		}
	}

	return details, fmt.Errorf("%W", "No matching electronics found")
}

func (e electronicsDetails) getSelections(choices []electronicsDetails) (selections []string) {
	selections = make([]string, 0)
	for _, nextSelection := range choices {
		selections = append(selections, nextSelection.description)
	}

	return
}

func (e electronicsDetails) getTons() float32 {
	return e.massByTL[techDetails.offset] + computerMass[electronics.computer]
}

func (e electronicsDetails) getCost() float32 {
	return e.costByTL[techDetails.offset] + computerCost[electronics.computer]
}

func (e electronicsDetails) getelectronicsCrew() (int, string) {
	return 0, ""
}
