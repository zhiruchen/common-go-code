package server

// chat rooms
type channel struct {
	name    string
	clients map[*client]bool
}

func newChannel(chanName string) *channel {
	return &channel{
		name:    chanName,
		clients: make(map[*client]bool, 8),
	}
}

func (c *channel) broadcast(s string, m []byte) {
	msg := append([]byte(s), ": "...)
	msg = append(msg, m...)
	msg = append(msg, '\n')

	for cl := range c.clients {
		cl.conn.Write(msg)
	}
}
