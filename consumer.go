package main

import (
	"github.com/streadway/amqp"

	"code.google.com/p/goprotobuf/proto"
	"github.com/Sophiacom/RabbitMQGoDemo/protobuf"
)

func Consume(connection *amqp.Connection, queue string, outChannel chan *protobuf.Notification) {
	// create a channel on this connection
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// start consuming data
	consumerTag := queue + "-consumer"
	deliveries, err := channel.Consume(
		queue,       // name
		consumerTag, // consumerTag
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		panic(err)
	}
	defer channel.Cancel(consumerTag, false)

	for delivery := range deliveries {
		notif := &protobuf.Notification{}
		proto.Unmarshal(delivery.Body, notif)

		outChannel <- notif

		if err := delivery.Ack(false); err != nil {
			panic(err)
		}
	}
}
