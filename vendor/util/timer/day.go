package timer

import (
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

func TimeZone() int64 {
	_, t := time.Now().Zone()
	return int64(t)
}

type Day struct {
	t time.Time
}

func NewDay(t time.Time) Day {
	return Day{DayStart(t)}
}

func DayStart(t time.Time) time.Time {
	yy, mm, dd := t.Date()
	t = time.Date(yy, mm, dd, 0, 0, 0, 0, t.Location())
	return t
}

const dayFormat = "2006-01-02"

func (d Day) String() string {
	return d.t.Format(dayFormat)
}

func (d Day) Begin() time.Time {
	return d.t
}

func (d Day) UnixBegin() int64 {
	return d.t.Unix()
}

func (d Day) Diff(d2 Day) int {
	return int(d.t.Sub(d2.t).Hours()) / 24
}

func (d Day) End() time.Time {
	return d.Begin().Add(time.Hour*24 - 1)
}

func (d Day) UnixEnd() int64 {
	return d.End().Unix()
}

func Today() Day {
	return NewDay(time.Now())
}
