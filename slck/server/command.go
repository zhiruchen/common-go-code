package server

type ID int

const (
	REG ID = iota
	JOIN
	LEAVE
	MSG
	CHNS
	USRS
)

type command struct {
	id       ID     // command type
	receiver string // the receiver of the command, a @user or a #channel
	sender   string // sender of the command, a @user
	body     []byte // message body
}
