package packets

import (
	"net"
)

type SendPacket struct{}

func (p SendPacket) Params() int {
	return 2
}

func (p SendPacket) VerifyParameters(params []string) bool {
	return true
}

func (p SendPacket) Handle(params []string, c net.Conn) {
	return
}

func (p SendPacket) Identifier() int {
	return 3
}
