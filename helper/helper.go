package helper

import (
	"crypto/md5"
	"fmt"
)

func MD5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
