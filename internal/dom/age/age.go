package age

import (
	"fmt"
	"strings"
	"time"
)

const AgeIntent = "AgeIntent"

var nameToAge = map[string]string{
	"bob": "1994-02-26T08:30:00Z",
}

func GetTotalAgeFromName(name string) string {
	dob, ok := nameToAge[strings.ToLower(name)]
	if !ok {
		return "person or object undefined"
	}
	dateOfbirth, _ := time.Parse(time.RFC3339, dob)

	years, months, days, hours, mins, secs := diff(dateOfbirth, time.Now())
	return fmt.Sprintf(
		"%s is %d years, %d months, %d days, %d hours, %d minutes and %d seconds old!",
		name,
		years,
		months,
		days,
		hours,
		mins,
		secs,
	)
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
