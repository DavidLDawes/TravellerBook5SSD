package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	// setup the bits for the view - see view.go for definitions
	settings = widget.NewForm()
	secondSettings = widget.NewForm()
	details = widget.NewVBox()

	a := app.New()
	w := a.NewWindow("Traveller SRD Capital Ship Star Ship Designer")

	// tech first since other things depend on it
	techDetails.init(settings, details)
	hullDetails.init(settings, details)
	drives.init(settings, details)
	electronics.init(secondSettings, details)
	// weapons.init(settings, details)
	// Always do hull & drives (not to mention tech) before berths
	// vehicles.initsecondSettings, details)
	berths.init(secondSettings, details)
	summary.init(settings, details)

	ui := widget.NewHBox(settings, secondSettings, details)
	w.SetContent(ui)
	w.ShowAndRun()
}

func getTons() (result float32) {
	result = hullDetails.getTons() + techDetails.getTons() + drives.getTons() +
		electronics.getTons() + berths.getTons()

	return
}

func getCargo() (result float32) {
	result = float32(hullDetails.tons) - getTons()

	return
}

func getCost() (result float32) {
	result = hullDetails.getCost() + drives.getCost() +
		electronics.getCost() + berths.getCost()

	return
}
