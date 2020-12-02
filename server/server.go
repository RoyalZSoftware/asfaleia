package main

import (
	"fmt"
	lib "github.com/royalzsoftware/asfaleia/lib"
	"net"
)

func handleConnection(c net.Conn) {
	fmt.Println(c)
}

func main() {
	sf := lib.ConstructPacket(0, []string{"PublicKey", "OtherShit"})
	fmt.Println(sf)
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
