package helpers

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = len(alphabet)
)

func Base62Encode(number uint64) string {
	var encoded strings.Builder
	encoded.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encoded.WriteByte(alphabet[number%uint64(length)])
	}

	return encoded.String()

}

func GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func CompareHashPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
