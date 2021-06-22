package server

import "fmt"

type hub struct {
	channels       map[string]*channel
	clients        map[string]*client
	cmds           chan command
	deregisterChan chan *client
	registerCh     chan *client
}

func newHub() *hub {
	return &hub{
		channels:       make(map[string]*channel),
		clients:        make(map[string]*client),
		cmds:           make(chan command, 1),
		deregisterChan: make(chan *client, 1),
		registerCh:     make(chan *client, 1),
	}
}

func (h *hub) run() {
	for {
		select {
		case client := <-h.registerCh:
			h.register(client)
		case client := <-h.deregisterChan:
			h.unregister(client)
		case cmd := <-h.cmds:
			switch cmd.id {
			case JOIN:
				h.joinChannel(cmd.sender, cmd.receiver)
			case LEAVE:
				h.leaveChannel(cmd.sender, cmd.receiver)
			case MSG:
				h.message(cmd.sender, cmd.receiver, cmd.body)
			default:
				h.sendBack(cmd.sender, []byte("not supported cmd"))
				fmt.Println("not supported cmd")
			}
		}
	}
}

func (h *hub) register(cli *client) {
	if _, ok := h.clients[cli.userName]; ok {
		cli.userName = ""
		cli.conn.Write([]byte("ERR username taken\n"))
		return
	}

	h.clients[cli.userName] = cli
	cli.conn.Write([]byte("OK\n"))
}

func (h *hub) unregister(cli *client) {
	if _, ok := h.clients[cli.userName]; ok {
		delete(h.clients, cli.userName)

		for _, ch := range h.channels {
			delete(ch.clients, cli)
		}
	}
}

func (h *hub) joinChannel(userName, chanName string) {
	if cli, ok := h.clients[userName]; ok {
		if ch, ok := h.channels[chanName]; ok {
			ch.clients[cli] = true
			return
		}

		newChan := newChannel(chanName)
		newChan.clients[cli] = true
		h.channels[chanName] = newChan
	}
}

func (h *hub) leaveChannel(userName, chanName string) {
	if cli, ok := h.clients[userName]; ok {
		if ch, chExists := h.channels[chanName]; chExists {
			delete(ch.clients, cli)
		}
	}
}

func (h *hub) message(userName, receiver string, msg []byte) {
	sender, ok := h.clients[userName]
	if !ok {
		return
	}

	switch receiver[0] {
	case '#':
		if ch, ok := h.channels[receiver[1:]]; ok {
			if _, cliExists := ch.clients[sender]; cliExists {
				ch.broadcast(sender.userName, msg)
			}
		}
	case '@':
		if to, ok := h.clients[receiver[1:]]; ok {
			to.conn.Write(append(msg, '\n'))
		}

	default:
		sender.conn.Write([]byte("receiver not found\n"))
	}
}

func (h *hub) sendBack(userName string, msg []byte) {
	if cli, ok := h.clients[userName]; ok {
		cli.conn.Write(append(msg, '\n'))
	}
}
