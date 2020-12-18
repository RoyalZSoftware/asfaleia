package packets

import (
	"fmt"
	"github.com/royalzsoftware/asfaleia/lib"
	"github.com/royalzsoftware/asfaleia/lib/utils"
	"math/rand"
	"net"
)

func generateRandomValidationMessage(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

/*
Type: Client-Packet
Usage: 0;<PUB-KEY>
Returns: 1;<HANDLE>;<ENCRYPTED-SYMMETRIC-KEY>
*/
type AuthPacket struct{}

func (p AuthPacket) Params() int {
	return 1
}

func (p AuthPacket) VerifyParameters(params []string) bool {
	keyValid := utils.VerifyInput(params[0], utils.RuleSet{
		MinLength: 128,
		MaxLength: 128,
	})
	return keyValid
}

func (p AuthPacket) Handle(params []string, c net.Conn) {
	wrongUsage := &WrongUsagePacket{}

	encrypter := &utils.RSAEncrypter{}
	err, publicKey := utils.BytesToPublicKey([]byte(params[0]))

	if err != nil {
		lib.SendAnswerPacket(wrongUsage, c)
		return
	}

	validationMessage := utils.EncodeInBase64(
		generateRandomValidationMessage(10),
	)
	fmt.Println(string(validationMessage))
	handle := len(lib.HandleQueue)
	err, encryptedMessage := encrypter.Encrypt(
		[]byte(validationMessage),
		publicKey,
	)

	var answer lib.ResponsePacket
	if err != nil {
		answer = wrongUsage
	} else {
		answer = &AuthVerifyIdentityPacket{
			Handle:           handle,
			EncryptedMessage: utils.EncodeInBase64(encryptedMessage),
		}
		lib.HandleQueue[handle] = string(validationMessage)
	}

	lib.SendAnswerPacket(answer, c)
}

func (p AuthPacket) Identifier() int {
	return 0
}
