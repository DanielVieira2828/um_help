package cryptoutil

import (
	"crypto/ed25519"
	"crypto/hmac"
	"encoding/hex"

	"github.com/DanielVieirass/um_help/config"
	"golang.org/x/crypto/sha3"
)

type Cryptoutil struct {
	config     *config.Config
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

func (c *Cryptoutil) HashPassword(str string) string {
	hasher := hmac.New(sha3.New256, []byte(c.config.CryptoConfig.HS256Password))
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
