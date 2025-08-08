package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// 1. 连接 Redis 服务器
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.228.128:6379", // Redis 服务器地址
		Password: "",                     // Redis 密码（如有请填写）
		DB:       0,                      // 使用的 Redis 数据库编号
	})

	// 2. 连接 MySQL 服务器
	// dsn 格式：用户名:密码@tcp(IP:端口)/数据库名?参数
	dsn := "chenhh:123456@tcp(192.168.228.128:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("MySQL连接失败:", err)
	}
	defer db.Close()

	key := "user:1:name" // Redis 缓存的 key

	// 3. 先从 Redis 查询数据
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Redis 未命中，去 MySQL 查询
		fmt.Println("Redis未命中，查MySQL...")
		// 假设 MySQL 有 user 表，字段 id, name
		var name string
		err = db.QueryRow("SELECT name FROM user WHERE id = ?", 1).Scan(&name)
		if err != nil {
			log.Fatal("MySQL查询失败:", err)
		}
		// 查询到后写入 Redis，设置缓存有效期 1 小时
		err = rdb.Set(ctx, key, name, time.Hour).Err()
		if err != nil {
			log.Fatal("写入Redis失败:", err)
		}
		fmt.Println("从MySQL获取并写入Redis，name:", name)
	} else if err != nil {
		// Redis 查询出错
		log.Fatal("Redis查询失败:", err)
	} else {
		// Redis 命中，直接返回
		fmt.Println("从Redis获取，name:", val)
	}
}
