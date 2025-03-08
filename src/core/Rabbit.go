package core

import (
  "log"
  amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}


func ConnectionRabbitMQ() *amqp.Connection {
	connStr := "amqps://iqgdawgs:xTsm02QZzRyXPjL3H7qhpdXCdxonFe5d@whale.rmq.cloudamqp.com/iqgdawgs"
	conn, err := amqp.Dial(connStr)
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}