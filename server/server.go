package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/royalzsoftware/asfaleia/lib"
	"github.com/royalzsoftware/asfaleia/lib/packets"
	"net"
	"strconv"
	"strings"
)

func handleConnection(c net.Conn) {
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if err := HandlePacket(string(data), c); err != nil {
			c.Write([]byte(err.Error()))
		}
	}
}

func main() {
	l, err := net.Listen("tcp4", "0.0.0.0:3308")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func HandlePacket(packetString string, c net.Conn) error {
	msgParts := strings.Split(packetString, ";")
	identifier, err := strconv.Atoi(msgParts[0])
	params := msgParts[1:]
	if err != nil {
		return err
	}
	packet := GetPacketByIdentifier(identifier)
	if packet == nil {
		return errors.New("packet: Packet not found")
	}
	if len(params) != packet.Params() || !packet.VerifyParameters(params) {
		return errors.New("packet: Wrong usage.")
	}
	packet.Handle(params, c)
	return nil
}

func GetPacketByIdentifier(identifier int) lib.RequestPacket {
	packetList := []lib.RequestPacket{
		&packets.AuthPacket{},
		&packets.GetPacket{},
		&packets.SendPacket{},
	}

	for _, packet := range packetList {
		if packet.Identifier() == identifier {
			return packet
		}
	}
	return nil
}
