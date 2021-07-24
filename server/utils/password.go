package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))
	return hashedPW, nil
}

func ComparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")
	fmt.Println(pwsalt)
	fmt.Println(HashPassword("123456"))
	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, fmt.Errorf("无效密码")
	}
	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	return hex.EncodeToString(shash) == pwsalt[0], nil
}

func Md5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
