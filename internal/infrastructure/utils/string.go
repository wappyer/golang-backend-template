package utils

import (
	"math/rand"
	"strings"
)

func RandStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func UnderscoreToLowercaseCamelCase(s string) string {
	if s == "" {
		return s
	}
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s[:1]) + s[1:] // 首字母小写
	return s
}
