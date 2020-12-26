package packets

import (
	"github.com/royalzsoftware/asfaleia/lib"
	"github.com/royalzsoftware/asfaleia/lib/utils"
	"math/rand"
	"net"
)

func generateEncodedSymmetricKey(size int) string {
	token := make([]byte, size)
	rand.Read(token)
	return utils.EncodeInBase64(token)
}

func encryptSymmetricKey(
	symmetricKey string,
	pubKey string,
) (error, []byte) {
	encrypter := &utils.RSAEncrypter{}
	err, publicKey := utils.BytesToPublicKey([]byte(pubKey))

	if err != nil {
		return err, []byte("")
	}

	err, encryptedKey := encrypter.Encrypt(
		[]byte(symmetricKey),
		publicKey,
	)

	if err != nil {
		return err, []byte("")
	}

	return nil, encryptedKey
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
	symmetricKey string,
	pubKey string,
	handle int,
) lib.ResponsePacket {
	err, encryptedSymmetricKey := encryptSymmetricKey(
		symmetricKey,
		pubKey,
	)
	if err != nil {
		return &WrongUsagePacket{}
	} else {
		answer := &AuthVerifyIdentityPacket{
			Handle:           handle,
			EncryptedMessage: utils.EncodeInBase64(encryptedSymmetricKey),
		}

		return answer
	}

}

func (p AuthPacket) Handle(params []string, c net.Conn) {

	encodedSymmetricKey := generateEncodedSymmetricKey(10)

	responsePacket := p.TryCreateVerifyPacket(
		encodedSymmetricKey,
		params[0],
		len(lib.HandleQueue),
	)
	lib.SendAnswerPacket(responsePacket, c)
}

func (p AuthPacket) Identifier() int {
	return 0
}
