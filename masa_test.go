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
