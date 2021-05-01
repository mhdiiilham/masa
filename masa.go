package masa

import (
	"regexp"
	"time"
)

type Waktu struct {
	time.Time
	Format  string
	Pattern []string
}

func New(time time.Time) Masa {
	return &Waktu{
		Time: time,
	}
}

func (w *Waktu) Convert(format string) Masa {
	w.Format = format
	w.getPattern()
	return w
}

func (w *Waktu) String() string {
	for _, pattern := range w.Pattern {
		format := string(pattern)
		locale := w.locale(format)
		re := regexp.MustCompile(`((dddd+|ddd+|dd+|DD+|D+|MMMM+|MMM+|MM+|M+|YYYY+|YY+))`)
		w.Format = string(re.ReplaceAll([]byte(w.Format), []byte(locale)))
	}
	return w.Format
}

func (w *Waktu) getPattern() {
	re := regexp.MustCompile(`((dddd|ddd|dd|DD|D|MMMM|MMM|MM|M|YYYY|YY))`)
	formatsBytes := re.FindAll([]byte(w.Format), -1)
	for _, fb := range formatsBytes {
		w.Pattern = append(w.Pattern, string(fb))
	}
}

func (w *Waktu) locale(format string) string {
	switch format {
	case "dddd":
		return fullDays(w.Time)
	case "ddd":
		return shortDays(w.Time)
	case "dd":
		return minDays(w.Time)
	default:
		return ""
	}
}
