package hash

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5加密
func ToMd5(code string) string {
	w := md5.New()
	io.WriteString(w, code)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}
