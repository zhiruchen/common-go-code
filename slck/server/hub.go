package server

type hub struct {
	channels       map[string]*channel
	clients        map[string]*client
	cmds           chan command
	deregisterChan chan *client
	registerCh     chan *client
}
