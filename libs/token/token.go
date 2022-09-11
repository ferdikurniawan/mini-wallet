package token

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/hex"
)

func GenSalt(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func GenToken(withSalt string) string {
	sha := sha1.New()
	sha.Write([]byte(withSalt))
	token := hex.EncodeToString(sha.Sum(nil))
	return token
}
