package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

// drive structure defines what we need to know about a given drive
type drive struct {
	tons float32 // tonnage used by the engine
	cost float32 // cost of the engine, in MCr
	perf int     // Actual performance, like -2 or -3. So a Jump drive J-2 has 2 for perf.
	eng  int     // Required minimum number of engineers
}

// drivedetails gathers drives, fuel, and the output UI stuff needed externally
// Most of the methods in this module are off of this driveDetails struct.
type driveDetails struct {
	j    drive
	m    drive
	p    drive
	fuel float32
	anti bool
}

var (
	drives = driveDetails{
		j:    drive{6, 12, 2, 1},
		m:    drive{2.5, 1.3, 2, 0},
		p:    drive{4, 2, 2, 0},
		fuel: 43,
		anti: false,
	}

	driveSelections = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}
	powerSelections = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}

	tech4JDrives = []int{0, 1, 2, 3, 4, 5, 6, 7}

	jSelect *widget.Select
	mSelect *widget.Select
	pSelect *widget.Select
	amCheck *widget.Check

	driveForm = widget.NewForm(
		widget.NewFormItem("Jump Drive", jSelect),
		widget.NewFormItem("Maneuver Drive", mSelect),
		widget.NewFormItem("Power Plant", pSelect),
		widget.NewFormItem("Antimatter", amCheck))

	driveLabel      = *widget.NewLabel("Drive details go here")
	driveDetailsBox = widget.NewVBox(&driveLabel)

	// Jump drives cost 2 MCr/ton
	jCost = float32(2)
	// Maneuvewr drives cost .5 MCr/ton
	mCost = float32(0.5)

	// Fraction of ship tonnage required, indexed by drive # (0-8 for jump & maneuver as they are optional)
	mTon = []float32{0.0, 0.01, 0.0125, 0.015, 0.0175, 0.025, 0.0325, 0.04, 0.0475}
	jTon = []float32{0.0, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.075, 0.08}
	// Note: pTon starts with a value for P-1, since you MUST include power.
	pTon       = []float32{0.015, 0.02, 0.025, 0.03, 0.04, 0.05, 0.065, 0.08}
	pTonByTech = []float32{1.25, 1.0, 1.0, 1.0, 1.0, 0.75, 0.6, 0.5}
)

func (d driveDetails) init(form *widget.Form, box *widget.Box) {
	driveLabel.SetText(d.getDriveDetails())
	box.Children = append(box.Children, driveDetailsBox)

	jSelect = widget.NewSelect(driveSelections, stringValuedNothing)
	jSelect.SetSelected("2")
	jSelect.Selected = "2"
	jSelect.OnChanged = d.jChange
	jSelect.Show()

	mSelect = widget.NewSelect(driveSelections, stringValuedNothing)
	mSelect.SetSelected("2")
	mSelect.Selected = "2"
	mSelect.OnChanged = d.mChange
	mSelect.Show()

	pSelect = widget.NewSelect(powerSelections, stringValuedNothing)
	pSelect.SetSelected("2")
	pSelect.Selected = "2"
	pSelect.OnChanged = d.pChange
	pSelect.Show()

	amCheck = widget.NewCheck("Antimatter", d.amChecked)
	amCheck.Checked = false
	amCheck.Show()

	form.AppendItem(widget.NewFormItem("Jump Drive", jSelect))
	form.AppendItem(widget.NewFormItem("Maneuver Drive", mSelect))
	form.AppendItem(widget.NewFormItem("Power Plant", pSelect))
	form.AppendItem(widget.NewFormItem("Antimatter", amCheck))
}

func (d driveDetails) jChange(newJump string) {
	if len(newJump) > 0 && len(newJump) < 3 {
		for offset, nextValue := range driveSelections {
			if newJump == nextValue {
				d.j.perf = offset
				d.j.tons = jTon[offset] * float32(hullDetails.tons)
				d.updateDrives()

				break
			}
		}
	}
}

func (d driveDetails) mChange(newManeuver string) {
	if len(newManeuver) > 0 && len(newManeuver) < 3 {
		for offset, nextValue := range driveSelections {
			if newManeuver == nextValue {
				d.m.perf = offset
				d.m.tons = mTon[offset] * float32(hullDetails.tons)
				d.updateDrives()

				break
			}
		}
	}
}

func (d driveDetails) pChange(newPower string) {
	if len(newPower) > 0 && len(newPower) < 3 {
		for offset, nextValue := range powerSelections {
			if newPower == nextValue {
				d.p.perf = offset
				d.p.tons = pTon[d.p.perf-1] * hullDetails.getHullTons()
				// d.p.cost =
				// d.p.eng =
				d.updateDrives()

				break
			}
		}
	}
}

func (d driveDetails) amChecked(antimatter bool) {
	d.anti = antimatter
	d.updateDrives()
}

func (d driveDetails) getDriveDetails() (driveDetails string) {
	if d.j.perf > 0 {
		driveDetails = fmt.Sprintf("Jump Drive J-%d, %.1f tons costing %.1f MCr\n",
			d.j.perf, d.j.tons, d.j.cost)
	}
	if d.m.perf > 0 {
		driveDetails += fmt.Sprintf("Maneuver Drive M-%d, %.1f tons costing %.1f MCr\n",
			d.m.perf, d.m.tons, d.m.cost)
	}
	driveDetails += fmt.Sprintf("Power Plant P-%d, %.1f tons costing %.1f MCr\n",
		d.p.perf, d.p.tons, d.p.cost)

	if d.anti {
		driveDetails += fmt.Sprintf("Antimatter stacks %.1f tons costing %.1f MCr\n",
			d.fuel/10, d.fuel*2.0)
	} else {
		driveDetails += fmt.Sprintf("Fuel %.1f tons costing %.1f MCr\n",
			d.fuel, d.fuel*0.1)
	}

	return
}

func (d driveDetails) updateDrives() (jChanged bool, mChanged bool) {
	jChanged = false
	maxJ := 8
	maxM := 8
	// maxP := 20

	switch techDetails.offset {
	case 0:
		maxJ = 0
		maxM = 5
	case 1:
		maxJ = 1
		maxM = 6
	case 2:
		maxJ = 2
		maxM = 6
	case 3:
		maxJ = 3
		maxM = 6
	case 4:
		maxJ = 4
		maxM = 6
	case 5:
		maxJ = 5
		maxM = 6
	case 6:
		maxJ = 6
		maxM = 7
	case 7:
		maxJ = 7
		maxM = 8
	}
	if maxJ > d.p.perf {
		maxJ = d.p.perf
	}
	if d.j.perf > maxJ {
		jChanged = true
		jSelect.Selected = strconv.Itoa(maxJ)
		d.j.perf = maxJ
	}

	if maxM > d.p.perf {
		maxM = d.p.perf
	}
	if d.m.perf > maxM {
		mChanged = true
		mSelect.Selected = strconv.Itoa(maxM)
		d.m.perf = maxM
	}
	driveLabel.SetText(d.getDriveDetails())

	return
}

func (d driveDetails) figureTechEffects(tlMin int) (techCost float32, techDiscount float32, techFuelDiscount float32,
	techPowerFactor float32) {
	techDiscount = 1.0
	techCost = 1.0
	techFuelDiscount = 1.0
	techPowerFactor = 1.0
	switch techDetails.offset {
	default:
	case 0:
		techPowerFactor = 1.25
		techFuelDiscount = 0.9
		break
	case 1:
		techFuelDiscount = 0.9
		break
	case 2:
		techCost = 1.25
		techDiscount = 0.5
		techFuelDiscount = 0.9
	case 3:
		techCost = 1.5
		techDiscount = 0.33333
		techFuelDiscount = 0.8
	case 4:
		techCost = 2.0
		techDiscount = 0.25
		techFuelDiscount = 0.7
	case 5:
		techCost = 2.0
		techDiscount = 0.2
		techFuelDiscount = 0.62
		techPowerFactor = 0.75
	case 6:
		techCost = 2.2
		techDiscount = 0.166
		techFuelDiscount = 0.55
	case 7:
		techCost = 2.3
		techDiscount = 0.14
		techFuelDiscount = 0.5
	case 8:
		techCost = 2.3
		techDiscount = 0.13
		techFuelDiscount = 0.45
	case 9:
		techCost = 2.3
		techDiscount = 0.125
		techFuelDiscount = 0.4
	case 10:
		techCost = 2.4
		techDiscount = 0.1111
		techFuelDiscount = 0.3333
	case 11:
		techCost = 2.4
		techDiscount = 0.1
		techFuelDiscount = 0.3
	case 12:
		techCost = 2.5
		techDiscount = 0.0909
		techFuelDiscount = 0.25
	case 13:
		techCost = 2.5
		techDiscount = 0.08181
		techFuelDiscount = 0.2222
	case 14:
		techCost = 2.6
		techDiscount = 0.075
		techFuelDiscount = 0.2
	}

	return
}

func stringValuedNothing(_ string) {
}

func (d driveDetails) getTons() float32 {
	return d.getJDriveTons() + d.getMDriveTons() + d.getPPlantTons() + d.getFuelTons()
}

func (d driveDetails) getCost() float32 {
	return d.j.cost + d.m.cost + d.p.cost
}

func (d driveDetails) getJDriveTons() float32 {
	return d.j.tons
}

func (d driveDetails) getMDriveTons() float32 {
	return d.m.tons
}

func (d driveDetails) getPPlantTons() float32 {
	return d.p.tons
}

func (d driveDetails) getFuelTons() float32 {
	return d.fuel
}

func getCrew() (count int, description string) {
	count = engineersFromDriveTonnage(drives.j.tons) + engineersFromDriveTonnage(drives.p.tons) +
		engineersFromDriveTonnage(drives.m.tons)

	description = fmt.Sprintf("%dxJump Engineers, %dxManeuver Engineers, %dxPower Engineers\n",
		engineersFromDriveTonnage(drives.j.tons), engineersFromDriveTonnage(drives.m.tons),
		engineersFromDriveTonnage(drives.p.tons))
	return
}

func engineersFromDriveTonnage(tonnage float32) int {
	if tonnage == 0.0 {
		return 0
	} else {
		switch techDetails.offset {
		case 5:
			return int(.99999 + tonnage/200.0)
		case 6:
			return int(.99999 + tonnage/250.0)
		case 7:
			return int(.99999 + tonnage/333.33333)
		case 8:
			return int(.99999 + tonnage/425.0)
		case 12:
			return int(.99999 + tonnage/500.0)
		case 13:
			return int(.99999 + tonnage/500.0)
		case 14:
			return int(.99999 + tonnage/500.0)
		default:
			return int(.99999 + tonnage/100.0)
		}
	}
}

func (d driveDetails) getIndexFromDrive(dString string) int {
	for resultInt, dMatch := range TrvIndex {
		if dMatch == dString {
			return resultInt
		}
	}

	return -1
}

func (d driveDetails) getDriveFromIndex(index int) string {
	return TrvIndex[index]
}
