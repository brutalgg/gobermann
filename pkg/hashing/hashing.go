package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

//Md5Hash a
func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
