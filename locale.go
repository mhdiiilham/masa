package masa

import "time"

func fullDays(time time.Time) string {
	weekdays := map[string]string{
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunda":     "Minggu",
	}
	return weekdays[time.Weekday().String()]
}

func shortDays(time time.Time) string {
	weekdays := map[string]string{
		"Monday":    "Sen",
		"Tuesday":   "Sel",
		"Wednesday": "Rab",
		"Thursday":  "Kam",
		"Friday":    "Jum",
		"Saturday":  "Sab",
		"Sunda":     "Min",
	}
	return weekdays[time.Weekday().String()]
}

func minDays(time time.Time) string {
	weekdays := map[string]string{
		"Monday":    "Sn",
		"Tuesday":   "Sl",
		"Wednesday": "Rb",
		"Thursday":  "Km",
		"Friday":    "Jm",
		"Saturday":  "Sb",
		"Sunda":     "Mn",
	}
	return weekdays[time.Weekday().String()]
}
