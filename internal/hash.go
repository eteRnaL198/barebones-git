package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func CalcHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	hashInHex := hex.EncodeToString(hashInBytes)

	return hashInHex, nil
}
