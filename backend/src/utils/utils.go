package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}
