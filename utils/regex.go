package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetIntByRegex(text, key string) (int, bool) {
	re, _ := regexp.Compile(fmt.Sprintf(" %s=\"([\\d]+)", key))
	match := re.FindStringSubmatch(text)

	if len(match) < 2 {
		return 0, false
	}

	i, _ := strconv.Atoi(match[1])

	return i, true
}

func GetStrByRegex(text, key string) (string, bool) {
	re, _ := regexp.Compile(fmt.Sprintf(` %s=\"(.*?)\"`, key))
	match := re.FindStringSubmatch(text)

	if len(match) < 2 {
		return "", false
	}

	return match[1], true
}

func GetStrDataByRegex(text, key string) (string, bool) {
	re, _ := regexp.Compile(fmt.Sprintf("%s><!\\[CDATA\\[(.*?)\\]", key))
	match := re.FindStringSubmatch(text)

	if len(match) < 2 {
		return "", false
	}

	return match[1], true
}

func GetXmlByRegex(text, key string) (string, bool) {
	re, _ := regexp.Compile(fmt.Sprintf("<%s>(.*?)</%s>", key, key))
	match := re.FindStringSubmatch(text)

	if len(match) < 2 {
		return "", false
	}

	return match[1], true
}

func GetStrQuotByRegex(text, key string) (string, bool) {
	re, _ := regexp.Compile(fmt.Sprintf(` %s=&quot;(.*?)&quot;`, key))
	match := re.FindStringSubmatch(text)

	if len(match) < 2 {
		return "", false
	}

	return match[1], true
}
