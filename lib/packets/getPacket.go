package packets

import (
	// lib "github.com/royalzsoftware/asfaleia/lib"
	"net"
)

type GetPacket struct{}

func (p GetPacket) Params() int {
	return 1
}

func (p GetPacket) VerifyParameters(params []string) bool {
	return true
}

func (p GetPacket) Handle(params []string, c net.Conn) {
	return
}

func (p GetPacket) Identifier() int {
	return 4
}
