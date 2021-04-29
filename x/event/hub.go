package event

import (
	"sync"
)

const (
	SmallHubLen  = 64
	MediumHubLen = 512
	LargeHubLen  = 1024
)

type Hub struct {
	mutex sync.RWMutex
	lines map[*Line]*Line
	len   int
}

func NewHub(channelLength int) *Hub {
	return &Hub{
		len:   channelLength,
		lines: make(map[*Line]*Line),
	}
}

func (h *Hub) disconnect(l *Line) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	delete(h.lines, l)
}

func (hub *Hub) Emit(v interface{}) {
	hub.mutex.RLock()
	defer hub.mutex.RUnlock()
	for _, line := range hub.lines {
		line.emit(v)
	}
}

func (hub *Hub) NewLine() *Line {
	var line = newLine(hub.len)
	line.Connect(hub)
	return line
}
