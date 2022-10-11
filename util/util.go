package util

import (
	"math/rand"
	"regexp"
	"time"
)

// RandomString 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	var result = make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// RandomNumber 生成随机数字编号
func RandomNumber(n int) string {
	var letters = []byte("1234567890")
	var result = make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// CheckEmail 检查邮箱的合法性
func CheckEmail(email string) bool {
	if matched, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !matched {
		return false
	}
	return true
}

// CheckMobile 检查手机的合法性
func CheckMobile(mobile string) bool {
	const regular = "^(13[0-9]|14[57]|15[0-35-9]|18[0-9]|19[9])\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

