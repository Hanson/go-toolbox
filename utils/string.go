package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func SubStr(s string, max int) string {
	sr := []rune(s)
	if len(sr) > max {
		sr = sr[:max]
	}
	return string(sr)
}

const (
	RandomStringModNumberPlusLetter = iota
	RandomStringModNumberPlusLetterPlusSymbol
	RandomStringModNumber
)

var (
	numberPlusLetter       = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberPlusLetterSymbol = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*(){}|\\[]?/")
	numberOnly             = []byte("0123456789")
)

func RandStr(length int, mod uint32) string {
	var key []byte

	switch mod {
	case RandomStringModNumberPlusLetter:
		key = numberPlusLetter
	case RandomStringModNumberPlusLetterPlusSymbol:
		key = numberPlusLetterSymbol
	case RandomStringModNumber:
		key = numberOnly
	default:
		key = numberPlusLetter
	}

	// 使用全局随机源，避免频繁创建
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = key[rand.Intn(len(key))]
	}

	return string(bytes)
}

func IsNumeric(str string, trim bool) bool {
	if trim {
		str = strings.TrimSpace(str)
	}
	_, err := strconv.Atoi(str)
	return err == nil
}

// 移除连着的换行符，最多保留一个，目前用于AI场景
func RemoveNewline(str string) string {
	re := regexp.MustCompile(`\n{2,}`)
	return re.ReplaceAllString(str, "\n")
}
