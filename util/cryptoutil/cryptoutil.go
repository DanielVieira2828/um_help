package cryptoutil

import (
	"crypto/ed25519"
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/consts"
	"github.com/go-jose/go-jose/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ssh"
)

type Cryptoutil struct {
	config     *config.Config
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

func parsePublicKey(key string) (ed25519.PublicKey, error) {
	raw, _, _, _, err := ssh.ParseAuthorizedKey([]byte(key))
	if err != nil {
		return nil, err
	}

	switch crypto := raw.(type) {
	case ssh.CryptoPublicKey:
		switch publicKey := crypto.CryptoPublicKey().(type) {
		case ed25519.PublicKey:
			return publicKey, nil
		default:
			return nil, errors.New("[ed25519] failed to parse public key")
		}
	default:
		return nil, errors.New("[ssh] failed to parse authorized key")
	}
}

func parsePrivateKey(key string, password string) (ed25519.PrivateKey, error) {
	privateKey, err := ssh.ParseRawPrivateKeyWithPassphrase([]byte(key), []byte(password))
	if err != nil {
		return nil, err
	}

	switch privateKey := privateKey.(type) {
	case *ed25519.PrivateKey:
		return *privateKey, nil
	default:
		return nil, errors.New("[ed25519] failed to parse private key")
	}
}

func New(cfg *config.Config) (*Cryptoutil, error) {
	publicKey, err := parsePublicKey(cfg.CryptoConfig.JWSPublicKey)
	if err != nil {
		return nil, err
	}

	privateKey, err := parsePrivateKey(cfg.CryptoConfig.JWSPrivateKey, cfg.CryptoConfig.JWSPrivateKeyPassword)
	if err != nil {
		return nil, err
	}

	return &Cryptoutil{
		config:     cfg,
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

func (c *Cryptoutil) signClaims(claims interface{}) (string, error) {

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.EdDSA, Key: c.privateKey}, nil)
	if err != nil {
		return "", err
	}

	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	jws, err := signer.Sign(payload)
	if err != nil {
		return "", err
	}

	result, err := jws.CompactSerialize()
	if err != nil {
		return "", err
	}

	return result, nil
}

type TokenClaims struct {
	SignId         string `json:"jti"`
	Issuer         string `json:"iss"`
	IssuedAt       int64  `json:"iat"`
	Subject        int64  `json:"sub"`
	ExpirationTime int64  `json:"exp"`
	Type           string `json:"typ"`
}

type RefreshTokenClaims struct {
	SignId string `json:"jwi"`
	Type   string `json:"typ"`
}

type SignResult struct {
	SignId         string `json:"sign_id"`
	JWS            string `json:"jws"`
	ExpirationTime int64  `json:"exp"`
	RefreshToken   string `json:"refresh_token"`
}

func (c *Cryptoutil) SignUser(userID int64) (*SignResult, error) {
	now := time.Now().Unix()
	expirationTime := time.Now().Add(time.Hour * time.Duration(c.config.CryptoConfig.JWSExpirationTimeInHours)).Unix()

	claims := &TokenClaims{
		SignId:         uuid.New().String(),
		Issuer:         c.config.InternalConfig.ServiceName,
		IssuedAt:       now,
		Subject:        userID,
		ExpirationTime: expirationTime,
		Type:           consts.AcessTokenType,
	}

	jws, err := c.signClaims(claims)
	if err != nil {
		return nil, err
	}

	refreshClaims := &RefreshTokenClaims{
		SignId: claims.SignId,
		Type:   consts.RefreshTokenType,
	}

	refreshToken, err := c.signClaims(refreshClaims)
	if err != nil {
		return nil, err
	}

	return &SignResult{
		SignId:         claims.SignId,
		JWS:            jws,
		ExpirationTime: expirationTime,
		RefreshToken:   refreshToken,
	}, nil
}

func (c *Cryptoutil) VerifyJWS(jws string, value interface{}) error {
	object, err := jose.ParseSigned(jws, []jose.SignatureAlgorithm{jose.EdDSA})
	if err != nil {
		return err
	}

	payload, err := object.Verify(c.publicKey)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(payload, value); err != nil {
		return err
	}

	return nil
}

func (c *Cryptoutil) HashString(str string) string {
	hasher := hmac.New(sha3.New256, []byte(c.config.CryptoConfig.HS256Password))
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
