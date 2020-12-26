package packets

import (
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

func generateRandomEncodedValidationMessage(size int) string {
	return utils.EncodeInBase64(generateRandomValidationMessage(size))
}

func generateFinishedValidationMessage(validationMessage string,
	pubKey string) (error, []byte) {
	encrypter := &utils.RSAEncrypter{}
	err, publicKey := utils.BytesToPublicKey([]byte(pubKey))

	if err != nil {
		return err, []byte("")
	}

	err, encryptedMessage := encrypter.Encrypt(
		[]byte(validationMessage),
		publicKey,
	)

	if err != nil {
		return err, []byte("")
	}

	return nil, encryptedMessage
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
		MinLength: 218,
		MaxLength: 218,
	})
	return keyValid
}

func (p AuthPacket) Handle(params []string, c net.Conn) {
	validationMessage := generateRandomEncodedValidationMessage(10)
	err, encryptedMessage := generateFinishedValidationMessage(
		validationMessage,
		params[0],
	)
	if err != nil {
		lib.SendAnswerPacket(&WrongUsagePacket{}, c)
	} else {
		handle := len(lib.HandleQueue)
		lib.HandleQueue[handle] = string(validationMessage)

		answer := &AuthVerifyIdentityPacket{
			Handle:           handle,
			EncryptedMessage: utils.EncodeInBase64(encryptedMessage),
		}

		lib.SendAnswerPacket(answer, c)
	}

}

func (p AuthPacket) Identifier() int {
	return 0
}
