package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

type berthDetails struct {
	description string
	count       int
	tonsPer     float32 // Tons per occupant
	costPer     float32 // Cost per opccupant
	military    bool    // Only on military?
	capital     bool    // Only on capital ships?
	small       bool    // Available on small ships?
}

type allBerths struct {
	berths []*berthDetails
}

var (
	sum = widget.NewLabel("Summary")

	lowLevel = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
		"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
	}
	medLevel = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
		"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
		"40", "41", "42", "43", "44", "45", "46", "47", "48", "49",
		"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
		"60", "61", "62", "63", "64", "65", "66", "67", "68", "69",
		"70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "40",
		"80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
		"90", "91", "92", "93", "94", "95", "96", "97", "98", "99",
		"100", "101", "102", "103", "104", "105", "106", "107", "108", "109",
		"110", "111", "112", "113", "114", "115", "116", "117", "118", "119",
		"120",
	}
	berthDetailsBox = widget.NewVBox(sum)
	berthTons       float32
	berthCost       float32

	milBarracks          = berthDetails{"Barracks", 0, 0.5, 0.5, true, false, true}
	twoFer               = berthDetails{"Double Bunks", 0, 1.5, 1.0, true, true, true}
	stateroom            = berthDetails{"Staterooms", 0, 4.0, 1.0, true, true, true}
	staffStateroom       = berthDetails{"Staff Staterooms", 3, 3.0, 1.0, true, true, true}
	luxuryStateRoom      = berthDetails{"Luxury Staterooms", 0, 5.0, 2.0, false, false, true}
	staffLuxuryStateRoom = berthDetails{"Staff Luxury Staterooms", 0, 5.0, 2.0, false, false, true}
	xStateroom           = berthDetails{"Xeno Staterooms", 0, 6.0, 10.0, false, true, false}
	luxuryXStateRoom     = berthDetails{"Luxury Xeno Staterooms", 0, 7.0, 12.0, false, true, false}
	residence            = berthDetails{"Residence", 0, 12.0, 8.0, false, false, false}
	xResidence           = berthDetails{"Xeno Residence", 0, 15.0, 14.0, false, false, false}
	luxurySuite          = berthDetails{"Luxury Suite", 0, 10.0, 6.0, false, false, false}
	luxuryXSuite         = berthDetails{"Luxury Xeno Suite", 0, 12.0, 9.0, false, false, false}
	low                  = berthDetails{"Low Berth", 0, 0.5, 0.05, false, false, false}

	berths = allBerths{
		[]*berthDetails{
			&milBarracks, &twoFer, &stateroom, &staffStateroom, &luxuryStateRoom, &staffLuxuryStateRoom, &xStateroom,
			&residence, &xResidence, &luxuryXStateRoom, &luxurySuite, &luxuryXSuite, &low,
		},
	}

	/*
		staffBerths = []berthDetails{twoFer, staffStateroom, low}

		milBerths   = []berthDetails{milBarracks, twoFer, staffStateroom, low}
		// add low berths to passengerBerths ?
		passengerBerths = []berthDetails{stateroom, xStateroom, luxuryStateRoom, luxuryXStateRoom, residence,
			xResidence, luxurySuite, luxuryXSuite, low}
	*/

	lowBerthSlider *widget.Slider
	// doubleSlider           *widget.Slider
	stateroomSlider *widget.Slider
	// xStateroomSlider       *widget.Slider
	// luxuryStateroomSlider  *widget.Slider
	// luxuryXStateroomSlider *widget.Slider
	// residenceSlider        *widget.Slider
	// xResidenceSlider       *widget.Slider
	// luxurySuiteSlider      *widget.Slider
	// luxuryXSuiteSlider     *widget.Slider
)

func (b allBerths) init(form *widget.Form, box *widget.Box) {
	stateroomSlider = widget.NewSlider(0.0, 200.0)
	stateroomSlider.Value = 0
	stateroomSlider.Show()
	stateroomSlider.OnChanged = stateroomsChanged

	lowBerthSlider = widget.NewSlider(0.0, 150.0)
	lowBerthSlider.Value = 0
	lowBerthSlider.Show()
	lowBerthSlider.OnChanged = lowBerthChanged

	form.AppendItem(widget.NewFormItem("Staterooms", stateroomSlider))
	form.AppendItem(widget.NewFormItem("Low Berths", lowBerthSlider))

	box.Children = append(box.Children, berthDetailsBox)
	updateBerths()
}

func (b allBerths) fixupSliders() {
	lowBerthSlider = widget.NewSlider(0.0, float64(hullDetails.tons)/5.0)
	stateroomSlider = widget.NewSlider(0.0, float64(hullDetails.tons)/1.25)
}

func stateroomsChanged(newSelection float64) {
	if stateroom.count != int(newSelection) {
		stateroom.count = int(newSelection)
		updateBerths()
	}
}

func lowBerthChanged(newSelection float64) {
	if low.count != int(newSelection) {
		low.count = int(newSelection)
		updateBerths()
	}
}

func updateBerths() {
	count := 0
	tons := float32(0.0)
	cost := float32(0.0)
	crew, summ := getOperationsCrew()
	staffStateroom.count = crew
	for _, berth := range berths.berths {
		count += berth.count
		tons += float32(berth.count) * berth.tonsPer
		cost += float32(berth.count) * berth.costPer
		if berth.count > 0 {
			summ += fmt.Sprintf(" %dx %s costing %.1f for %.1f tons\n",
				berth.count, berth.description, float32(berth.count)*berth.costPer,
				float32(berth.count)*berth.tonsPer)
		}
	}
	sum.SetText(fmt.Sprintf("%sTotal %d berths costing %.1f using %.1f tons",
		summ, count, cost, tons))
	sum.Refresh()
	sum.Show()
	berthTons = tons
	berthCost = cost

	summary.update()
}

func (b allBerths) getTons() float32 {
	return berthTons
}

func (b allBerths) getCost() float32 {
	return berthCost
}

func getBerthsCrew(staff2support int) (crew int, description string) {
	crew = 1 + (staff2support)/8
	description = fmt.Sprintf("%d stewards", crew)

	return
}

func getBerthsDetails(selection string) (details berthDetails) {
	for _, match := range berths.berths {
		if match.description == selection {
			details = *match

			return
		}
	}

	return
}

func getOperationsCrew() (int, string) {
	if hullDetails.tons < 1000 {
		// pilot, nav
		return 2, "Pilot, Navigator\n"
	} else if hullDetails.tons < 5000 {
		// purser/captain, comms, pilot, 2xnav

		return 5, "Purser/Captain, Comms, Pilot, 2xNavigator\n"
	} else if hullDetails.tons < 25000 {
		// purser, captain, comms, 2xpilot, 2xnav, 2xsecurity, support

		return 10, "Purser, Captain, Comms, 2xPilot, 2xNavigator, 2xSecurity, Support\n"
	} else if hullDetails.tons < 100000 {
		// purser, captain, 2xcomms, sensors, 2xpilot, 2xnav, 8xsecurity, 2xsupport

		return 19, "Purser, Captain, 2xComms, Sensors, 2xPilot, 2xNavigator, 8xSecurity, 2xSupport\n"
	} else if hullDetails.tons < 500000 {
		// purser, captain, 2xcomms, 2xsensors, 2xpilot, 2xnav, 12xsecurity, 8xsupport

		return 28, "Purser, Captain, 2xComms, 2xSensors, 2xPilot, 2xNavigator, 12xSecurity, 6xSupport\n"
	} else if hullDetails.tons < 2500000 {
		// purser, captain, 4xcomms, 2xsensors, 4xpilot, 4xnav, 20xsecurity, 10xsupport

		return 46, "Purser, Captain, 4xComms, 2xSensors, 4xPilot, 4xNavigator, 20xSecurity, 10xSupport\n"
	} else if hullDetails.tons < 5000000 {
		// purser, captain, 4x dept. heads, 8xcomms, 4xsensors, 4xpilot, 4xnav, 40xsecurity, 20xsupport, 10 maint

		return 96,
			"Purser, Captain, 4xDept. Heads, 8xComms, 4xSensors, 4xPilot, 4xNavigator, 40xSecurity, 20xSupport, " +
				"10 Maintenance\n"
	} else if hullDetails.tons < 10000000 {
		// purser, bursar, 2xbankers, commander, captain, 4xlt., 8x dept. heads, 12xcomms, 8xsensors,
		// 8xpilot, 4xnav, 80xsecurity, 40xsupport, 20 maint

		return 190,
			"Purser, Bursar, 2xBankers, Commander, Captain, 4xLt., 8xDept. Heads, 12xComms, 8xSensors, 8xPilot, " +
				"4xNavigator, 80xSecurity, 40xSupport, 20 Maintenance\n"
	} else if hullDetails.tons < 25000000 {
		// 4x pursers, 2xbursers, 4xbankers, commander, captain, 12xlt., 20x dept. heads, 12xcomms, 16xsensors,
		// 10xpilot, 6xnav, 120xsecurity, 60xsupport, 40 maint

		return 308,
			"4x Pursers, 2xBursars, 4xBankers, Commander, Captain, 12xLt., 20xDept. Heads, 12xComms, 16xSensors, " +
				"10xPilot, 6xNavigator, 120xSecurity, 60xSupport, 40 Maintenance\n"
	} else {
		// 8x pursers, 4xbursers, 8xbankers, 4xcommander, 4xcaptain, 36xlt., 24x dept. heads, 16xcomms,
		// 16xsensors, 12xpilot, 8xnav, 200xsecurity, 100xsupport, 50 maint
		return 490,
			"8x Pursers, 4xBursars, 8xBankers, 4xCommanders, 4xCaptains, 24xLt., 24xDept. Heads, 16xComms, 16xSensors, " +
				"12xPilot, 8xNavigator, 200xSecurity, 100xSupport, 40 Maintenance\n"
	}
}
