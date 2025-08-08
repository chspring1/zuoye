package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	password := "password1234"
	hashpwd := "b9c950640e1b3740e98acb93e669c65766f6670dd1609ba91ff41052ba48c6f3"
	md5pwd := "bdc87b9c894da5168059e00ebffb9077"
	hash := hashPassword(password)
	fmt.Printf("原始密码: %s\n", password)
	fmt.Printf("哈希密码: %s\n", hash)
	md5pawd := md5Hash(password)
	fmt.Printf("MD5密码: %s\n", md5pawd)
	hashPwd := verifyHash(password, hashpwd)
	fmt.Printf("哈希密码验证: %v\n", hashPwd)
	md5Pwd := verifyMD5Hash(password, md5pwd)
	fmt.Printf("MD5密码验证: %v\n", md5Pwd)
}

// 哈希加密实列
func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// md5加密实例
func md5Hash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// 哈希反向验证
func verifyHash(password, hash string) bool {
	return hashPassword(password) == hash
}

// MD5反向验证
func verifyMD5Hash(password, hash string) bool {
	return md5Hash(password) == hash
}
