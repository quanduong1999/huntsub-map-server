package timer

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDayStart(t *testing.T) {
	start := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)
	d := Day{t: DayStart(start.Add(time.Hour))}
	diff := d.Begin().Unix() - start.Unix()
	if diff != 0 {
		t.Logf("begin: %d", diff)
		t.Fatal("begin of day is not correct")
	}
}

func TestDayEnd(t *testing.T) {
	end := time.Date(2017, time.January, 1, 23, 59, 59, 999, time.Local)
	d := Day{t: DayStart(end.Add(-time.Hour))}
	diff := d.End().Unix() - end.Unix()
	if diff != 0 {
		t.Logf("begin: %d", diff)
		t.Fatal("begin of day is not correct")
	}
}
