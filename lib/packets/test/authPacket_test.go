package test

import (
	"github.com/matryer/is"
	"github.com/royalzsoftware/asfaleia/lib/packets"
	"testing"
)

func TestAuthPacketParameterVerification(t *testing.T) {
	is := is.New(t)
	packet := &packets.AuthPacket{}

	isValid := packet.VerifyParameters([]string{"Test"})

	is.True(!isValid)
}

func TestAuthPacketReturnsWrongUsage(t *testing.T) {
	is := is.New(t)
	packet := &packets.AuthPacket{}

	responsePacket := packet.TryCreateVerifyPacket(
		"SymmetricKey",
		"Test",
		0,
	)

	is.Equal(responsePacket.Identifier(), 102)
}

func TestAuthPacketReturnsEncryptedSymmetricKey(t *testing.T) {
	is := is.New(t)
	packet := &packets.AuthPacket{}

	publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC1DGPUS" +
		"PqPJdcWvdBNM5ymAQlKIXbLMa3X4D0JaBKKqv47w4ii3" +
		"sLw2lckEHk+8LGyBNkbQN0aaVxWeuo4R7IsBMXVpJgvf" +
		"/iQImRm5b9HIxkAYi5PqYGcHKBqiuzuHmpPlGYdGpM+h" +
		"K5EeWiDj/sRVcP42KyYzGmd6ExbqaWMewIDAQAB"

	responsePacket := packet.TryCreateVerifyPacket(
		"SymmetricKey",
		publicKey,
		0,
	)

	is.Equal(responsePacket.Identifier(), 101)
}
