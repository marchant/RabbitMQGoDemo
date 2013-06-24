package main

import (
	"math/rand"
	"time"

	"code.google.com/p/goprotobuf/proto"
	"github.com/Sophiacom/RabbitMQGoDemo/protobuf"
)

var (
	minTime uint32 = 100
	maxTime uint32 = 200
)

func Send(inChannel chan *protobuf.Notification, outChannel chan *protobuf.Log) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for notif := range inChannel {
		sendingTime := time.Duration(minTime+uint32(uint64(r.Uint32())*uint64(maxTime-minTime)>>32)) * time.Millisecond
		time.Sleep(sendingTime)

		outChannel <- &protobuf.Log{
			Notification: notif,
			SendingTime:  proto.Int32(int32(sendingTime)),
		}
	}
}
