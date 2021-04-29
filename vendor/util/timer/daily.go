package timer

import (
	"context"
	"time"
)

type DailyTimer struct {
	hour   int
	minute int
	last   time.Time
	C      chan time.Time
}

func NewDailyTimer(hour int, minute int) *DailyTimer {
	return &DailyTimer{
		hour:   hour,
		minute: minute,
		C:      make(chan time.Time, 16),
	}
}

func (t *DailyTimer) Start(ctx context.Context) {
	go t.launch(ctx)
}

func (timer *DailyTimer) launch(ctx context.Context) {
	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	ticker := time.NewTicker(time.Second * 6)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			if timer.isTimeToEmit(t) {
				timer.emit(t)
			}
		case <-localCtx.Done():
			return
		}
	}
}

func (timer *DailyTimer) emit(t time.Time) {
	timer.C <- t
	timer.last = t
}

func (timer *DailyTimer) isTimeToEmit(t time.Time) bool {
	// already emitted
	if timer.last.Day() == t.Day() {
		return false
	}
	hh, mm, _ := t.Clock()
	return hh == timer.hour && mm == timer.minute
}
