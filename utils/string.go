package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func RandStr(length int64, mod uint32) string {
	var strKey string
	if mod == RandomStringModNumberPlusLetter {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if mod == RandomStringModNumberPlusLetterPlusSymbol {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*(){}|\\[]?/"
	} else if mod == RandomStringModNumber {
		strKey = "0123456789"
	} else {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = strKey[r.Intn(len(strKey))]
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
