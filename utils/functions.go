package utils

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// StringToUint64 将字符串转换成64位
func StringToUint64(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)

	return i
}

// GenerateCode 生成短信验证码
func GenerateCode() string {
	var code string

	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		number := rand.Intn(10)

		code += strconv.Itoa(number)
	}

	return code
}

// AnalysisHtml 解析html转换成string
func AnalysisHtml(fileName string,data interface{}) string {
	file,_ := template.ParseFiles(fileName)

	body := new(bytes.Buffer)

	file.Execute(body,data)

	return body.String()
}

// HashAndSalt 给密码加密
func HashAndSalt(password string) string  {
	hash,_ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)

	return string(hash)
}

// ComparePasswords 给密码解密
func ComparePasswords(hash string, password []byte) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, password)

	if err != nil {
		return false
	}

	return true
}

// TrimHtml 删除字符串中的html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}