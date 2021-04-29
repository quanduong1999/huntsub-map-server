package event

type Line struct {
	c    chan interface{}
	hubs []*Hub
}

func newLine(len int) *Line {
	return &Line{
		c:    make(chan interface{}, len),
		hubs: make([]*Hub, 0),
	}
}

func (l Line) C() <-chan interface{} {
	return l.c
}

func (line *Line) Connect(hub *Hub) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	hub.lines[line] = line
	line.hubs = append(line.hubs, hub)
}

func (line *Line) MultipleConnect(hubs ...*Hub) {
	for _, h := range hubs {
		line.Connect(h)
	}
}

func (line *Line) Disconnect() {
	for _, h := range line.hubs {
		h.disconnect(line)
	}
}

func (line Line) emit(v interface{}) {
	if cap(line.c) > len(line.c) {
		line.c <- v
	}
}
