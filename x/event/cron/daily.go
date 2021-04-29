package cron

import "time"

type dailyTimer struct {
	Timer
	at int64
}

var day int64 = 24 * 3600

func NewDailyTimer(at string, handler func()) (*dailyTimer, error) {
	var t, err = time.ParseInLocation("15:04", at, time.Local)
	if err != nil {
		return nil, err
	}
	var d = &dailyTimer{
		Timer: Timer{handler: handler},
	}
	d.at = t.Unix()

	return d, nil
}

func (t *dailyTimer) delay() int64 {
	// t.at is at year 0000
	return (t.at-time.Now().Unix())%day + day
}

func (t *dailyTimer) Start() {
	var delay = time.Second * time.Duration(t.delay())
	t.Schedule(delay, 24*time.Hour)
}
