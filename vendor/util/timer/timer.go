package timer

import "time"

type Timer struct {
	timer   *time.Timer
	handler func()
}

func NewTimer(handler func()) *Timer {
	return &Timer{
		handler: handler,
	}
}

func (e *Timer) Schedule(delay time.Duration, loop time.Duration) {
	e.timer = time.AfterFunc(delay, func() {
		if loop > 0 {
			e.Schedule(loop, loop)
		}
		e.handler()
	})
}

func (e *Timer) Cancel() {
	e.timer.Stop()
}
