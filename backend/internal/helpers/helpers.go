package helpers

import (
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = len(alphabet)
)

func Base62Encode(number uint64) string {
	var encoded strings.Builder
	encoded.Grow(7)
	for ; number > 0; number = number / uint64(length) {
		encoded.WriteByte(alphabet[number%uint64(length)])
	}

	result := encoded.String()

	runes := []rune(result)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)[:7]

}
