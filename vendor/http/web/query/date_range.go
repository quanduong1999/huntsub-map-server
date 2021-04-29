package query

import (
	"http/web"
	"time"
	"util/timer"
)

const dateFormat = "2006-01-02"

type DateRange struct {
	start timer.Day
	end   timer.Day
}

func (d DateRange) String() string {
	return d.start.String() + ":" + d.end.String()
}

func (d *DateRange) CountDay() int {
	return d.end.Diff(d.start)
}

func (d *DateRange) Start() time.Time {
	return d.start.Begin()
}

func (d *DateRange) StartUnix() int64 {
	return d.Start().Unix()
}

func (d *DateRange) End() time.Time {
	return d.end.End()
}

func (d *DateRange) EndUnix() int64 {
	return d.End().Unix()
}

func (q Query) GetDay(key string) (*timer.Day, error) {
	value := q.Get(key)
	if value == "" {
		return nil, web.BadRequest("missing " + key)
	}
	ti, err := time.Parse(dateFormat, value)
	if err != nil {
		return nil, web.WrapBadRequest(err, "date format for "+key+" must be "+dateFormat)
	}
	d := timer.NewDay(ti)
	return &d, nil
}

func (q Query) MustGetDay(key string) *timer.Day {
	day, err := q.GetDay(key)
	if err != nil {
		panic(err)
	}
	return day
}

func (q Query) MustGetDateRange(startKey string, endKey string) *DateRange {
	start := q.MustGetDay(startKey)
	end := q.MustGetDay(endKey)
	return &DateRange{start: *start, end: *end}
}
