package packets

import "strconv"

type AuthVerifyIdentityPacket struct {
	Handle           int
	EncryptedMessage string
}

func (p AuthVerifyIdentityPacket) Identifier() int {
	return 101
}

func (p AuthVerifyIdentityPacket) ToString() string {
	return strconv.Itoa(p.Handle) + ";" + p.EncryptedMessage
}
