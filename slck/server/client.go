package server

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

type client struct {
	conn       net.Conn
	outbound   chan<- command // connection between client and hub, use to send commands to hub.
	register   chan<- *client // client let hub know it want to register itself.
	deregister chan<- *client // let hub know user closed socket, hub should deregister the client.
	userName   string         // the user name of the user that sitting behind the TCP.
}

func (cli *client) read() error {
	for {
		msg, err := bufio.NewReader(cli.conn).ReadBytes('\n')
		if err == io.EOF { // connection closed, deregister this client from hub.
			cli.deregister <- cli
			return nil
		}
		if err != nil {
			return err
		}

		cli.handle(msg)
	}
}

func (cli *client) handle(msg []byte) {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(msg, []byte(" "))[0]))
	params := bytes.TrimSpace(bytes.TrimPrefix(msg, cmd))

	var err error
	switch string(cmd) {
	case "REG":
		err = cli.reg(params)
	case "JOIN":
		err = cli.join(params)
	case "LEAVE":
		err = cli.leave(params)
	case "MSG":
		err = cli.msg(params)
	case "CHNS":
		cli.chns()
	case "USRS":
		cli.usrs()
	default:
		cli.err(fmt.Errorf("unknow cmd: %s", cmd))
	}

	if err != nil {
		cli.err(err)
	}
}

func (cli *client) reg(params []byte) error {

}

func (cli *client) err(err interface{}) {

}

func (cli *client) join(params []byte) error {

}

func (cli *client) leave(params []byte) error {

}

func (cli *client) msg(params []byte) error {

}

func (cli *client) chns() {

}

func (cli *client) usrs() {

}
