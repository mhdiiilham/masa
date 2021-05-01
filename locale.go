package masa

import (
	"fmt"
	"time"
)

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

func fullMonths(time time.Time) string {
	weekdays := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}
	return weekdays[time.Month().String()]
}

func shortMonths(time time.Time) string {
	weekdays := map[string]string{
		"January":   "Jan",
		"February":  "Feb",
		"March":     "Mar",
		"April":     "Apr",
		"May":       "Mei",
		"June":      "Jun",
		"July":      "Jul",
		"August":    "Ags",
		"September": "Sep",
		"October":   "Okt",
		"November":  "Nov",
		"December":  "Des",
	}
	return weekdays[time.Month().String()]
}

func twoDigitMonth(time time.Time) string {
	return fmt.Sprintf("%02d", int(time.Month()))
}
func oneDigitMonth(time time.Time) string {
	return fmt.Sprintf("%d", int(time.Month()))
}
