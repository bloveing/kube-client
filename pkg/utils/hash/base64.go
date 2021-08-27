package hash

import (
	"encoding/base64"
)

// 加密base64
func EncodeBase64(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

// 解密base64
func DecodeBase64(sEnc string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(sEnc)
}
