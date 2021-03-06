package utils

import (
	"crypto/md5"
	"fmt"
	c "yeetikuserver/config"
)

// 加密密码，md5(md5(password + salt) + public_salt)
func EncryptPassword(password, salt string) string {
	saltedPassword := fmt.Sprintf("%x", md5.Sum([]byte(password))) + salt
	md5SaltedPassword := fmt.Sprintf("%x", md5.Sum([]byte(saltedPassword)))
	return fmt.Sprintf("%x", md5.Sum([]byte(md5SaltedPassword+c.GetConfig().PublicSalt)))
}
