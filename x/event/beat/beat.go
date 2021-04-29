package beat

import (
	"time"
)

var lastBeat *time.Time
var activated bool

func everySecond(v *time.Time) {
	if v.Second() == 0 {
		everyMinute(v)
	}
	lastBeat = v
}

func everyMinute(v *time.Time) {
	var dateFormat = "2006-01-02"
	if v.Format(dateFormat) != lastBeat.Format(dateFormat) {
		daily.Emit(v)
	}
}

func activate() {
	if activated {
		return
	}
	go loop()
	activated = true
}

func loop() {
	var tick = time.NewTicker(time.Second)
	for {
		v := <-tick.C
		everySecond(&v)
	}
}

func init() {
	activate()
}
