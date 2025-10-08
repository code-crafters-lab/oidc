package oauth2

import (
	"crypto/rand"
	"crypto/sha256"
	mr "math/rand"
	"time"

	"github.com/zitadel/oidc/v3/pkg/crypto"
)

type CodeChallengeMethod string

const (
	CodeChallengeMethodPlain CodeChallengeMethod = "plain"
	CodeChallengeMethodS256  CodeChallengeMethod = "S256"
)

type GeneratorOptions = func(*pkce)

type PKCE interface {
	GenerateCodeChallenge(options ...GeneratorOptions) string
	VerifyCodeChallenge(codeVerifier string) bool
	Print()
}

type pkce struct {
	codeVerifierLength  int
	codeVerifier        string
	codeChallenge       string
	codeChallengeMethod CodeChallengeMethod
}

func (p *pkce) shaCodeChallenge(code string) string {
	return crypto.HashString(sha256.New(), code, false)
}

func (p *pkce) generateCodeVerifier() string {
	if p.codeVerifier == "" {
		letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~")
		codeVerifier := make([]byte, p.codeVerifierLength)
		n, err := rand.Read(codeVerifier)
		if err != nil || n != p.codeVerifierLength {
			return ""
		}
		for i, b := range codeVerifier {
			codeVerifier[i] = letters[b%byte(len(letters))]
		}
		p.codeVerifier = string(codeVerifier)
	}
	return p.codeVerifier
}

func (p *pkce) GenerateCodeChallenge(opts ...GeneratorOptions) string {
	for _, opt := range opts {
		opt(p)
	}
	if p.codeVerifierLength < 43 || p.codeVerifierLength > 128 {
		//throw `Expected a code verifier length between 43 and 128. Received ${len}.`;
		// todo 警告日志
		p.codeVerifierLength = 64
	}
	p.generateCodeVerifier()
	switch p.codeChallengeMethod {
	case CodeChallengeMethodPlain:
		p.codeChallenge = p.codeVerifier
	case CodeChallengeMethodS256:
		p.codeChallenge = p.shaCodeChallenge(p.codeVerifier)
	}
	return p.codeChallenge
}

func (p *pkce) VerifyCodeChallenge(codeVerifier string) bool {
	if p.codeChallengeMethod == CodeChallengeMethodS256 {
		return p.codeChallenge == p.shaCodeChallenge(codeVerifier)
	} else if p.codeChallengeMethod == CodeChallengeMethodPlain {
		return p.codeChallenge == codeVerifier
	}
	return false
}

func (p *pkce) Print() {
	println("codeVerifier:", p.codeVerifier)
	println("codeChallengeMethod:", p.codeChallengeMethod)
	println("codeChallenge:", p.codeChallenge)
}

func NewPKCE() PKCE {
	return &pkce{codeVerifierLength: 64, codeChallengeMethod: CodeChallengeMethodPlain}
}

func WithCodeVerifier(codeVerifier string) GeneratorOptions {
	return func(p *pkce) {
		p.codeVerifier = codeVerifier
	}
}

func WithCodeVerifierLength(len int) GeneratorOptions {
	return func(p *pkce) {
		p.codeVerifierLength = len
	}
}

func WithRandomCodeVerifier() GeneratorOptions {
	return func(p *pkce) {
		p.codeVerifierLength = randomInRange(43, 128)
	}
}

func randomInRange(min, max int) int {
	r := mr.New(mr.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func WithCodeChallengeMethodPlain() GeneratorOptions {
	return func(p *pkce) {
		p.codeChallengeMethod = CodeChallengeMethodPlain
	}
}

func WithCodeChallengeMethodS256() GeneratorOptions {
	return func(p *pkce) {
		p.codeChallengeMethod = CodeChallengeMethodS256
	}
}
