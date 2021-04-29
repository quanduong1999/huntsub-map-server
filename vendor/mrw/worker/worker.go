package worker

import (
	"huntsub/huntsub-map-server/x/mlog"
	"time"
)

var logger = mlog.NewTagLog("worker")

const (
	goldenRatio         float32 = 1.618
	reverseGoldenRation float32 = 1 / goldenRatio
	max                         = 1024
	jobChanLen                  = 1 << 16
)

type job func()

type worker struct {
	min   int
	jobs  chan job
	sig   signalLine
	stopC chan struct{}
}

func newWorker(min int) *worker {
	if min < 1 {
		min = 1
	}
	var w = &worker{
		min:   min,
		sig:   newSignalLine(),
		jobs:  make(chan job, jobChanLen),
		stopC: make(chan struct{}, 1),
	}

	return w
}

func (w *worker) start() {
	go w.auto()
}

func (w *worker) Add(j job) {
	w.jobs <- j
}

func (w *worker) stop() {
	w.stopC <- struct{}{}
}

func (w *worker) loop(stop <-chan struct{}) {
	for {
		select {
		case r := <-w.jobs:
			r()
		case <-stop:
			return
		}
	}
}

func (w *worker) auto() {
	sec := time.Tick(time.Second)
	for {
		select {
		case <-sec:
			w.adjust()
		case <-w.stopC:
			return
		}
	}
}

func (w *worker) adjust() {
	jc := len(w.jobs)
	if jc < 1 {
		w.less()
		return
	}
	ratio := float32(jc) / float32(w.sig.Len())
	if ratio > goldenRatio {
		w.more()
		logger.Infof(0, "active up to %d", w.sig.Len())
	} else if ratio < reverseGoldenRation {
		w.less()
		logger.Infof(0, "active down to %d", w.sig.Len())
	}
}

func (w *worker) more() {
	if w.sig.Len() > max {
		return
	}
	stop := w.sig.More()
	go w.loop(stop)
}

func (w worker) less() {
	if w.sig.Len() <= w.min {
		return
	}
	w.sig.Less()
}

func NewWorker() *worker {
	w := newWorker(1)
	w.start()
	return w
}
