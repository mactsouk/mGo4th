package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ producer")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("amqp.Dial():", err)
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("Go", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Queue:", q)

	message := "Writing to RabbitMQ!"
	err = ch.Publish("", "Go", false, false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(message)},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Message published to Queue!")
}
