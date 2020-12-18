package utils

type RuleSet struct {
	MinLength int
	MaxLength int
}

func VerifyInput(input string, rule RuleSet) bool {
	if len(input) < rule.MinLength || len(input) > rule.MaxLength {
		return false
	}
	return true
}
