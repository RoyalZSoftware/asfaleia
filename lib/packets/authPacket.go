package packets

import (
	"github.com/royalzsoftware/asfaleia/lib"
	"github.com/royalzsoftware/asfaleia/lib/utils"
	"math/rand"
	"net"
)

func generateRandomEncodedValidationMessage(size int) string {
	token := make([]byte, size)
	rand.Read(token)
	return utils.EncodeInBase64(token)
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

func (p AuthPacket) TryCreateVerifyPacket(
	validationMessage,
	pubKey string,
	handle int,
) lib.ResponsePacket {
	err, encryptedMessage := generateFinishedValidationMessage(
		validationMessage,
		pubKey,
	)
	if err != nil {
		return &WrongUsagePacket{}
	} else {
		answer := &AuthVerifyIdentityPacket{
			Handle:           handle,
			EncryptedMessage: utils.EncodeInBase64(encryptedMessage),
		}

		return answer
	}

}

func (p AuthPacket) Handle(params []string, c net.Conn) {

	validationMessage := generateRandomEncodedValidationMessage(10)

	responsePacket := p.TryCreateVerifyPacket(
		validationMessage,
		params[0],
		len(lib.HandleQueue),
	)
	lib.SendAnswerPacket(responsePacket, c)
}

func (p AuthPacket) Identifier() int {
	return 0
}
