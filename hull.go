package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type hull struct {
	tons        int
	tonString   string
	isArmored   bool
	armor       int
	armorString string
	armorTons   int
	sections    int
}

const (
	defaultHull        = 1000
	defaultHullString  = "1000"
	defaultSections    = 1
	defaultArmor       = false
	defaultArmorString = "None"
	defaultArmorValue  = 0
)

var (
	hullDetails = hull{
		defaultHull, defaultHullString,
		defaultArmor, 0, defaultArmorString, defaultArmorValue,
		defaultSections,
	}

	hullSelect  *widget.Select
	armorSelect *widget.Select

	hullLabel  = *widget.NewLabel("Hull: 1000 tons")
	armorLabel = *widget.NewLabel("Unarmored")

	armorTonCostByTech = []float32{
		0.025, 0.025, .0125, .0125, .0125, .0125, .00833333, .00833333, .00666666, .005, .004, .003, .002, .002,
	}
	maxArmorByTech = []int{
		8, 9, 10, 11, 12, 13, 14, 15, 24, 34, 36, 42, 46, 50,
	}

	hullSizes = []string{
		"100", "120", "150", "200", "300", "330", "400", "500", "600", "700", "800", "900", "980",
		"1000", "1200", "1500", "2000", "2500", "2600", "3000", "3500", "4000", "4500", "5000", "5500", "6000", "6400", "6500", "7000", "7200", "7500", "8000", "8500", "9000", "9500",
		"10000", "11000", "12000", "13000", "14000", "14400", "15000", "16000", "18000", "20000", "22000", "22500", "25000", "27000", "30000", "35000", "40000", "45000", "50000", "55000", "60000", "64000", "65000", "70000", "72000", "75000", "80000", "85000", "90000", "95000",
		"100000", "110000", "115000", "120000", "125000", "130000", "133000", "140000", "144000", "150000", "156000", "160000", "170000", "175000", "180000", "190000", "200000", "210000", "220000", "225000", "250000", "300000", "320000", "350000", "375000", "400000", "450000", "480000", "500000", "550000", "600000", "640000", "650000", "700000", "720000", "750000", "800000", "850000", "900000", "950000",
		"1000000", "1100000", "1150000", "1200000", "1250000", "1300000", "1330000", "1400000", "1440000", "1500000", "1560000", "1600000", "1700000", "1750000", "1800000", "1900000", "2000000", "2100000", "2200000", "2250000", "2500000", "3000000", "3200000", "3500000", "4000000", "4500000", "4800000", "4900000", "5000000", "5500000", "6000000", "6400000", "6500000", "7000000", "7200000", "7500000", "8000000", "8500000", "9000000", "9500000",
		"10000000", "11000000", "12000000", "12500000", "13000000", "13300000", "14000000", "14400000", "15000000", "15600000", "16000000", "17000000", "17500000", "18000000", "19000000", "20000000", "21000000", "22000000", "22500000", "25000000", "30000000", "32000000", "35000000", "39900000", "40000000", "45000000", "48000000", "50000000", "55000000", "60000000", "64000000", "65000000", "70000000", "72000000", "75000000", "80000000", "85000000", "90000000", "95000000",
		"100000000", "115000000",
	}
)

func (h hull) init(form *widget.Form, box *widget.Box) {
	hullSelect = widget.NewSelect(hullSizes, stringValuedNothing)
	hullSelect.SetSelected(defaultHullString)
	hullSelect.PlaceHolder = defaultHullString
	hullSelect.Selected = defaultHullString
	hullSelect.OnChanged = h.hullChanged
	hullSelect.Show()

	armorSelect = widget.NewSelect(h.getArmorRangeFromTech(), nothing)
	armorSelect.SetSelected(defaultArmorString)
	armorSelect.PlaceHolder = noneString
	armorSelect.Selected = noneString
	armorSelect.OnChanged = h.armorChanged
	armorSelect.Show()

	form.AppendItem(widget.NewFormItem("Hull", hullSelect))
	form.AppendItem(widget.NewFormItem("Armor", armorSelect))
	box.Children = append(box.Children, &hullLabel, &armorLabel)

	h.tons = defaultHull
	h.tonString = defaultHullString
	hullDetails.armorTons = 0
	h.armorString = defaultArmorString
	h.isArmored = false
	h.armor = 0

	h.updateHullDetails()
	hullLabel.Refresh()
	hullLabel.Show()
	armorLabel.Refresh()
	armorLabel.Show()
}

func (h hull) hullChanged(hullSelected string) {
	for _, nextTonnage := range hullSizes {
		if hullSelected == nextTonnage {
			hTons, err := strconv.Atoi(hullSelected)
			if err == nil {
				h.tons = hTons
				h.tonString = hullSelected
				h.updateHullDetails()
				drives.updateDrives()
				summary.update()
			}
			break
		}
	}
}

func (h hull) armorChanged(armorSelected string) {
	if armorSelected == "None" {
		hullDetails.armorString = armorSelected
		hullDetails.armorTons = 0
		hullDetails.isArmored = false
		hullDetails.armor = 0
		hullDetails.updateArmorDetails()
		summary.update()
	} else {
		newArmor, err := strconv.Atoi(armorSelected)
		if err == nil {
			hullDetails.armor = newArmor
			hullDetails.armorString = armorSelected
			hullDetails.isArmored = true
			hullDetails.armorTons = newArmor * 2
			hullDetails.updateArmorDetails()
			summary.update()
		}
	}
}

func (h hull) getArmorRangeFromTech() (available []string) {
	available = make([]string, 1)
	available[0] = "None"
	for i := 1; i < maxArmorByTech[techDetails.offset]; i++ {
		available = append(available, strconv.Itoa(i))
	}

	return
}

func (h hull) updateHullDetails() {
	armorSelect.Options = h.getArmorRangeFromTech()
	if h.getSections() == 1 {
		hullLabel.SetText(fmt.Sprintf("Hull tonnage: %s, single section with %d hull/structure",
			h.tonString, int(float32(h.tons)/50.0)))
	} else {
		hullLabel.SetText(fmt.Sprintf("Hull tonnage: %s, using %d sections with %d hull/structure each",
			h.tonString, h.getSections(), int(float32(h.tons)/(float32(h.getSections()*50.0)))))
	}
	hullLabel.Refresh()
	hullLabel.Show()
}

func (h hull) updateArmorDetails() {
	if h.armor < 1 {
		armorLabel.SetText("Unarmored")
	} else {
		armorLabel.SetText(fmt.Sprintf("Armor AV-%s using %d tons", h.armorString, hullDetails.armorTons))
	}
	armorLabel.Refresh()
	armorLabel.Show()
}

func (h hull) getHullTons() float32 {
	return float32(h.tons)
}

func (h hull) getArmorTons() float32 {
	if h.isArmored {
		return float32(hullDetails.armorTons) / 20.0
	} else {
		return 0
	}
}

func (h hull) getTons() float32 {
	return h.getArmorTons()
}

func (h hull) getCost() float32 {
	return h.getArmorTons()
}

func (h hull) isCapital() bool {
	return h.tons > 999
}

func (h hull) isSmall() bool {
	return h.tons < 500
}

func (h hull) getSections() (result int) {
	if h.tons < 2000 {
		result = 1
	} else if h.tons < 10000 {
		result = 2
	} else if h.tons < 40000 {
		result = 3
	} else if h.tons < 200000 {
		result = 4
	} else if h.tons < 700000 {
		result = 5
	} else if h.tons < 1500000 {
		result = 6
	} else if h.tons < 2500000 {
		result = 7
	} else if h.tons < 4000000 {
		result = 8
	} else if h.tons < 6000000 {
		result = 9
	} else if h.tons < 9000000 {
		result = 10
	} else if h.tons < 15000000 {
		result = 11
	} else if h.tons < 22000000 {
		result = 12
	} else if h.tons < 30000000 {
		result = 13
	} else if h.tons < 40000000 {
		result = 14
	} else {
		result = 15
	}

	return
}
