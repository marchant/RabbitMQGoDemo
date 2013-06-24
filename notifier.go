package main

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/Sophiacom/RabbitMQGoDemo/protobuf"
)

var (
	queues = [3]string{"ANDROID", "IOS", "WINDOWS_PHONE"}
	uri    = "amqp://wowodc:wowodc@localhost:5672/wowodc-rabbit-2"
)

func main() {
	// open a connection to localhost
	connection, err := amqp.Dial(uri)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	senderToLogger := make(chan *protobuf.Log)

	for _, queue := range queues {
		fmt.Println("Consuming on", queue, "queue.")
		consumerToSender := make(chan *protobuf.Notification)
		go Consume(connection, queue, consumerToSender)
		go Send(consumerToSender, senderToLogger)
	}
	Log(connection, senderToLogger)
}
