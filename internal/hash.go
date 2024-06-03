package internal

import (
	"crypto/sha1"
	"encoding/hex"
)

func CalcHash(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	hashInBytes := hasher.Sum(nil)
	hashInHex := hex.EncodeToString(hashInBytes)
	return hashInHex
}
