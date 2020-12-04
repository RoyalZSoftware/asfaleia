package packets

import (
	"github.com/royalzsoftware/asfaleia/lib"
	"github.com/royalzsoftware/asfaleia/lib/utils"
	"net"
)

type AuthPacket struct{}

func (p AuthPacket) Params() int {
	return 1
}

func (p AuthPacket) Handle(params []string, c net.Conn) {
	encrypter := &utils.MockEncrypter{}
	validationMessage := "Hello World"
	handle := len(lib.HandleQueue)
	answer := &AuthVerifyIdentityPacket{
		Handle:           handle,
		EncryptedMessage: encrypter.Encrypt(string(validationMessage)),
	}

	lib.HandleQueue[handle] = validationMessage
	lib.SendAnswerPacket(answer, c)
	return
}

func (p AuthPacket) Identifier() int {
	return 0
}
