package server

// chat rooms
type channel struct {
	name    string
	clients map[*client]bool
}

func (c *channel) broadcase(s string, m []byte) {
	msg := append([]byte(s), ": "...)
	msg = append(msg, m...)
	msg = append(msg, '\n')

	for cl := range c.clients {
		cl.conn.Write(msg)
	}
}
