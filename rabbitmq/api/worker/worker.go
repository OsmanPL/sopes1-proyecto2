package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KromDaniel/rejonson"
	"github.com/go-redis/redis"
	"github.com/sevenpok/api-rabbit/controller"
	"github.com/sevenpok/api-rabbit/models"
	"github.com/streadway/amqp"
)

var goRedisClient = redis.NewClient(&redis.Options{
	Addr: "34.121.89.39:6379",
})

func insertRedis(body string) {
	client := rejonson.ExtendClient(goRedisClient)
	client.JsonArrAppend("squid_game", ".", body)
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	chDelivery, err := ch.Consume(
		"game",
		"",
		true,
		false,
		false,
		false, nil)

	if err != nil {
		log.Fatal(err)
	}

	noStop := make(chan bool)

	go func() {
		for delivery := range chDelivery {
			var game models.Game
			json.Unmarshal([]byte(delivery.Body), &game)
			//fmt.Println(game)
			err := controller.Create(game)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Se inserto correctamente en MongoDB")
			}
			insertRedis(string(delivery.Body))
		}
	}()

	<-noStop
}
