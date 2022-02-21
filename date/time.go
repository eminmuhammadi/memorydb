package date

import (
	"time"
	_ "time/tzdata"
)

// Return time in defined timezone using IANA timezone data
func Now(tz string) time.Time {
	localTime, err := time.LoadLocation(tz)

	if err != nil {
		panic(err)
	}

	return time.Now().In(localTime)
}
