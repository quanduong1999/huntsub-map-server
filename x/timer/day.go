package timer

import (
	"fmt"
	"strings"
	"time"
)

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * 60
	secondsPerDay    = 24 * secondsPerHour
	secondsPerWeek   = 7 * secondsPerDay
	daysPer400Years  = 365*400 + 97
	daysPer100Years  = 365*100 + 24
	daysPer4Years    = 365*4 + 1
)

var offset int64

func init() {
	var _, d = time.Now().Zone()
	offset = int64(d)
}

func LocalStartOfToday() int64 {
	var times int64
	var nowGTM = time.Now().Unix()
	var nowLocal = time.Now().Unix() - offset
	if TimeToDay(nowGTM) != TimeToDay(nowLocal) {
		if offset > 0 {
			var ellapsed = nowGTM % secondsPerDay
			times = nowGTM - ((ellapsed + offset) % secondsPerDay) + secondsPerDay
		} else {
			var ellapsed = nowGTM % secondsPerDay
			times = nowGTM - ((ellapsed + offset) % secondsPerDay) - secondsPerDay
		}

	} else {
		var ellapsed = nowGTM % secondsPerDay
		times = nowGTM - ((ellapsed + offset) % secondsPerDay)
	}
	return times

}

func TimeZone() int64 {
	_, t := time.Now().Zone()
	return int64(t)
}
func TimeToDay(ctime int64) string {
	var t = time.Unix(ctime, 0)
	return t.Format("2006-01-02")
}
func FormatDate(value string) string {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	var ch = strings.Split(t.String(), " ")
	return ch[0]
}

func TimeToDayUnix(ctime int64) string {
	var t = time.Unix(ctime, 0)
	return t.Format("2006-01-02T15:04:05Z07:00")
}
