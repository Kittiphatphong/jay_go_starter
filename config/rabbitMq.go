package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"go_starter/logs"
)

func InitRabbit() (*amqp.Connection, error) {

	myUser := GetEnv("amqp.username", "guest")
	myPassword := GetEnv("amqp.password", "guest")
	myHost := GetEnv("amqp.host", "localhost")
	myPort := GetEnv("amqp.port", "5672")

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", myUser, myPassword, myHost, myPort)
	fmt.Println("Rabbit dsn:", dsn)

	connectRabbitMQ, err := amqp.Dial(dsn)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	fmt.Println("Connected to RabbitMq")
	return connectRabbitMQ, nil
}
