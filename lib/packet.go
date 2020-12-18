package lib

import (
	"net"
	"strconv"
	"strings"
)

var HandleQueue = make(map[int]string)

type Packet interface {
	Identifier() int
}

type ResponsePacket interface {
	Identifier() int
	ToString() string
}

type RequestPacket interface {
	Identifier() int
	Params() int
	Handle(params []string, c net.Conn)
	VerifyParameters(params []string) bool
}

func SendAnswerPacket(packet ResponsePacket, c net.Conn) {
	c.Write([]byte(strconv.Itoa(packet.Identifier()) + ";" + packet.ToString()))
}

func ConstructPacket(identifier int, params []string) string {
	return strconv.Itoa(identifier) + ";" + strings.Join(params, ";")
}
