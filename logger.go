package main

import (
	"github.com/streadway/amqp"

	"code.google.com/p/goprotobuf/proto"
	"github.com/Sophiacom/RabbitMQGoDemo/protobuf"
)

var (
	exchange   = "logger-exchange"
	routingKey = ""
)

func Log(connection *amqp.Connection, inChannel chan *protobuf.Log) {
	// create a channel on this connection
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	for log := range inChannel {
		body, err := proto.Marshal(log)
		if err != nil {
			panic(err)
		}

		err = channel.Publish(
			exchange,   // publish to an exchange
			routingKey, // routing to 0 or more queues
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "",
				ContentEncoding: "",
				Body:            body,
				DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
				Priority:        0,              // 0-9
				// a bunch of application/implementation-specific fields
			},
		)
		if err != nil {
			panic(err)
		}
	}
}
