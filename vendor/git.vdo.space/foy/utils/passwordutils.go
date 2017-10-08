package utils

import (
	"crypto/sha1"
	"fmt"
)

//GetFinalPwdWithSalt 根据输入密码获取 最终密码和盐
func GetFinalPwdWithSalt(password string) (pwd string, salts string) {
	salt := RandomString(32)
	h := sha1.New()
	h.Write([]byte(salt + password))
	bs := fmt.Sprintf("%x", h.Sum(nil))
	return bs, salt
}

//GetFinalPwdBySalt 根据密码和盐获取最终密码
func GetFinalPwdBySalt(password string, salts string) (pwd string) {

	h := sha1.New()
	h.Write([]byte(salts + password))
	bs := fmt.Sprintf("%x", h.Sum(nil))

	return bs
}
