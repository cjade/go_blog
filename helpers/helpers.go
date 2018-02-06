package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(string(str)))
	strs := h.Sum(nil)
	return hex.EncodeToString(strs)
}
