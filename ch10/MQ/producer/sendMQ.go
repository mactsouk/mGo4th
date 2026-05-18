package main

import (
	"fmt"
	"context"   //imported to use context.Background() 
	amqp "github.com/rabbitmq/amqp091-go"
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
	
	//resolve nil pointer panic in sendMQ.go by using context.Background()
	err = ch.PublishWithContext(context.Background(), "", "Go", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Message published to Queue!")
}
