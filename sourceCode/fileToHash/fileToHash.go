package fileToHash

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
)

func CalcSHA256(filePath string) (string, error) {
	file_data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(file_data) // Hashing

	return hex.EncodeToString(hash[:]), nil
}
