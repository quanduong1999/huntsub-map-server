package event

import (
	"sync"
)

type Cancel func()
type Line chan interface{}
type HubLen int

const SmallHub = HubLen(16)
const MediumHub = HubLen(SmallHub * 16)
const LargeHub = HubLen(MediumHub * 16)

type Hub struct {
	sync.RWMutex
	m      map[Line]Cancel
	l      HubLen
	latest interface{}
}

func NewHub(l HubLen) *Hub {
	return &Hub{
		m: map[Line]Cancel{},
		l: l,
	}
}

func (h *Hub) NewLine() (Line, Cancel) {
	h.Lock()
	defer h.Unlock()
	c := Line(make(chan interface{}, h.l))
	h.m[c] = func() {}
	return Line(c), func() { h.Stop(c) }
}

func (h *Hub) Stop(l Line) {
	h.Lock()
	defer h.Unlock()
	delete(h.m, l)
}

func (h *Hub) Emit(v interface{}) {
	h.RLock()
	defer h.RUnlock()
	h.latest = v
	for l := range h.m {
		select {
		case l <- v:
		default:
		}
	}
}

func (h *Hub) Value() interface{} {
	return h.latest
}