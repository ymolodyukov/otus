package cryptor

import (
	"crypto/md5"
	"fmt"
)

func GetPasswordHash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
