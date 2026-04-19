package mq

import (
	"os"

	"github.com/nats-io/nats.go"
)

var (
	Nc *nats.Conn
	err error
)

func Init() {
	natsHost := os.Getenv("NATS_HOST")
	if natsHost == "" {
		natsHost = "localhost"
	}
	natsPort := os.Getenv("NATS_PORT")
	if natsPort == "" {
		natsPort = "4222"
	}
	url := "nats://" + natsHost + ":" + natsPort
	Nc, err = nats.Connect(url)
	if err != nil {
		panic(err)
	}
}