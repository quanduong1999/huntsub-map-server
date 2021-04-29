package worker

type signalLine []chan struct{}

func newSignalLine() signalLine {
	return signalLine(make([]chan struct{}, 0))
}

func (q *signalLine) More() chan struct{} {
	n := make(chan struct{}, 1)
	*q = append(*q, n)
	return n
}

func (q *signalLine) Less() {
	x := q.Len() - 1
	if x < 0 {
		return
	}
	n := (*q)[x]
	*q = (*q)[:x]
	n <- struct{}{}
}

func (q *signalLine) Len() int {
	return len(*q)
}
