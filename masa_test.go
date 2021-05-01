package masa

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMasa_Day(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "get full day name in Bahasa Indonesia",
			arg:  "dddd",
			want: "Jumat",
		},
		{
			name: "get short day name in Bahasa Indonesia",
			arg:  "ddd",
			want: "Jum",
		},
		{
			name: "get min day name in Bahasa Indonesia",
			arg:  "dd",
			want: "Jm",
		},
		{
			name: "argument not supported",
			arg:  "xxx",
			want: "xxx",
		},
		{
			name: "full day with prefix",
			arg:  "Hari ini adalah hari: dddd.",
			want: "Hari ini adalah hari: Jumat.",
		},
	}

	for _, tc := range testCases {
		goTime, _ := time.Parse("2006-01-02", "2000-12-15")
		day := New(goTime).Convert(tc.arg).String()
		assert.Equal(t, tc.want, day)
	}
}

func TestMasa_Month(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want string
		date string
	}{
		{
			name: "get full month name in Bahasa Indonesia",
			arg:  "MMMM",
			want: "Desember",
			date: "2000-12-15",
		},
		{
			name: "get short month name in Bahasa Indonesia",
			arg:  "MMM",
			want: "Des",
			date: "2000-12-15",
		},
		{
			name: "get 2 digit number month name in Bahasa Indonesia",
			arg:  "MM",
			want: "04",
			date: "1968-04-07",
		},
		{
			name: "get 1 digit number month name in Bahasa Indonesia",
			arg:  "M",
			want: "9",
			date: "1968-09-01",
		},
		{
			name: "Additional test case",
			arg:  "Indonesia merdeka di bulan: MMMM.",
			want: "Indonesia merdeka di bulan: Agustus.",
			date: "1945-08-17",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			goTime, _ := time.Parse("2006-01-02", tc.date)
			assert.Equal(t, tc.want, New(goTime).Convert(tc.arg).String())
		})
	}
}
