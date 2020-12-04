package packets

import (
	// "github.com/royalzsoftware/asfaleia/lib"
	"net"
)

type SendPacket struct{}

func (p SendPacket) Params() int {
	return 2
}

func (p SendPacket) Handle(params []string, c net.Conn) {
	return
}

func (p SendPacket) Identifier() int {
	return 3
}
