package lib

import (
	"net"
	"strconv"
	"strings"
)

type Packet interface {
	Params() int
	Handle(params []string, c net.Conn)
	Identifier() int
}

func ConstructPacket(identifier int, params []string) string {
	return strconv.Itoa(identifier) + ";" + strings.Join(params, ";")
}

func HandlePacket(packetString string, c net.Conn) error {
	msgParts := strings.Split(packetString, ";")
	identifier, err := strconv.Atoi(msgParts[0])
	params := msgParts[1:]
	if err != nil {
		return err
	}
	packets := []Packet{
		&AuthPacket{},
		&AnsPacket{},
		&SendPacket{},
		&GetPacket{},
	}
	for _, packet := range packets {
		if packet.Identifier() == identifier {
			packet.Handle(params, c)
		}
	}
	return nil
}

type AuthPacket struct{}

type AuthVerifyPacket struct{}

type AnsPacket struct{}

type SendPacket struct{}

type GetPacket struct{}

// Authpacket
func (p AuthPacket) Params() int {
	return 1
}

func (p AuthPacket) Handle(params []string, c net.Conn) {
	return
}

func (p AuthPacket) Identifier() int {
	return 0
}

// AuthVerify Packet
func (p AuthVerifyPacket) Params() int {
	return 1
}

func (p AuthVerifyPacket) Handle(params []string, c net.Conn) {
	return
}

func (p AuthVerifyPacket) Identifier() int {
	return 1
}

// Anspacket
func (p AnsPacket) Params() int {
	return 2
}

func (p AnsPacket) Handle(params []string, c net.Conn) {
	return
}

func (p AnsPacket) Identifier() int {
	return 2
}

// Sendpacket
func (p SendPacket) Params() int {
	return 2
}

func (p SendPacket) Handle(params []string, c net.Conn) {
	return
}

func (p SendPacket) Identifier() int {
	return 3
}

// Getpacket
func (p GetPacket) Params() int {
	return 1
}

func (p GetPacket) Handle(params []string, c net.Conn) {
	return
}

func (p GetPacket) Identifier() int {
	return 4
}
