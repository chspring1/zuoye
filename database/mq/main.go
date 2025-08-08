package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 生产者函数
func publish() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("连接RabbitMQ失败:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("打开通道失败:", err)
	}
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"test_queue", // 队列名
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal("声明队列失败:", err)
	}

	body := "Hello MQ!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal("发送消息失败:", err)
	}
	fmt.Println("消息已发送:", body)
}

// 消费者函数
func consume() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("连接RabbitMQ失败:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("打开通道失败:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test_queue", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("声明队列失败:", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("注册消费者失败:", err)
	}

	fmt.Println("等待接收消息...")
	for d := range msgs {
		fmt.Printf("收到消息: %s\n", d.Body)
		break // 只消费一条消息后退出
	}
}

func main() {
	// 先发送一条消息，再消费一条消息
	publish()
	consume()
}
