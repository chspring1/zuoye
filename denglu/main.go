package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

var (
	userData = make(map[string]string)
	mu       sync.RWMutex
)

func main() {
	// 示例：用户注册
	register("user1", "password123")

	// 示例：用户登录
	login("user1", "password123")
	login("user1", "wrongpassword")
}

// 注册用户
func register(username, password string) {
	mu.Lock()
	defer mu.Unlock()

	// 对密码进行哈希加密
	hash := sha256.New()
	hash.Write([]byte(password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))

	userData[username] = passwordHash
	fmt.Println("用户注册成功！")
	fmt.Printf("用户名: %s, 密码哈希: %s\n", username, passwordHash)
}

// 用户登录
func login(username, password string) {
	mu.RLock()
	defer mu.RUnlock()

	// 对输入的密码进行哈希加密
	hash := sha256.New()
	hash.Write([]byte(password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))

	if storedHash, ok := userData[username]; ok {
		if storedHash == passwordHash {
			fmt.Println("登录成功！")
			return
		}
	}
	fmt.Println("登录失败！")
}
