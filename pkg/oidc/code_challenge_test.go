package oidc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	codeVerifier  = "3GB.c0gBsluxqf4IO69M7.VO.EX2BVHhqzICXlpv_p6b8W-qwh.O9l4QDALthCmi"
	codeChallenge = "kARfceIQjgT9Sy1YNd1xKb97G4jf68pCuZuqbiNAg5g"
)

func TestVerifyCodeChallenge(t *testing.T) {
	tests := []struct {
		name          string
		codeChallenge *CodeChallenge
		codeVerifier  string
		expected      bool
	}{
		{"nil", nil, codeVerifier, false},
		{"plain", &CodeChallenge{Challenge: codeChallenge}, codeVerifier, false},
		{"S256", &CodeChallenge{Challenge: codeChallenge, Method: CodeChallengeMethodS256}, codeVerifier, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, VerifyCodeChallenge(tt.codeChallenge, tt.codeVerifier), "VerifyCodeChallenge(%v, %v)", tt.codeChallenge, tt.codeVerifier)
		})
	}
}
