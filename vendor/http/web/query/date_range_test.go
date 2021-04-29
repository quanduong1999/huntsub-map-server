package query

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateRangeUnmarshaler(t *testing.T) {
	sample := `{"start_date": "2017-01-01", "end_date": "2017-01-02"}`
	q := DateRange{}
	err := json.Unmarshal([]byte(sample), &q)
	if err != nil {
		t.Fatal(err.Error())
	}
	start, _ := time.Parse("2006-01-02", "2017-01-01")
	if q.StartUnix() != start.Unix() {
		t.Fatal("parse start date failed")
	}
	end, _ := time.Parse("2006-01-02", "2017-01-02")
	if q.EndUnix() != end.Unix() {
		t.Fatal("parse end date failed")
	}
}
