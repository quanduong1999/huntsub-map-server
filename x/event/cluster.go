package event

type Cluster struct {
	hubs map[string]*Hub
	any  *Hub
}

func NewCluster() *Cluster {
	return &Cluster{
		hubs: make(map[string]*Hub),
		any:  NewHub(MediumHubLen),
	}
}

func (c *Cluster) mustGetHub(name string) *Hub {
	var hub = c.hubs[name]
	if hub == nil {
		hub = NewHub(SmallHubLen)
		c.hubs[name] = hub
	}
	return hub
}

func (c *Cluster) OnAny(len int) *Line {
	if len < SmallHubLen {
		len = LargeHubLen
	}
	var line = newLine(len)
	line.Connect(c.any)
	return line
}

func (c *Cluster) Emit(name string, v interface{}) {
	c.any.Emit(v)
	var h = c.hubs[name]
	if h != nil {
		h.Emit(v)
	}
}
