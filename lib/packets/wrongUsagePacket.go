package packets

type WrongUsagePacket struct{}

func (p WrongUsagePacket) Identifier() int {
	return 102
}

func (p WrongUsagePacket) ToString() string {
	return "Wrong usage"
}
