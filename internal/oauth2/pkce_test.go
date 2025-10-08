package oauth2

import "testing"

func TestNewPKCE(t *testing.T) {
	p := NewPKCE()
	p.GenerateCodeChallenge(WithCodeChallengeMethodS256())
	p.Print()
}
