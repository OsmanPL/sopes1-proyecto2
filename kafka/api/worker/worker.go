package main

import(
	"github.com/segmentio/kafka-go"
	"context"
	"fmt"
	"log"
	"os"
	"github.com/KromDaniel/rejonson"
	"github.com/go-redis/redis"
)

const(
	topic = "Game"
	brokerAddress = "localhost:9092"
)
var goRedisClient = redis.NewClient(&redis.Options{
	Addr: "34.121.89.39:6379",
  })


func consume(ctx context.Context){
	l := log.New(os.Stdout,"Kafka escuchando: ",0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
		GroupID: "my-group",
		Logger: l,
	})

	for{
		msg, err := r.ReadMessage(ctx)
		if err != nil{
			panic("No se pudo leer el mensaje: "+ err.Error())
		}

		fmt.Println("Mensaje recibido: ", string(msg.Value))
		insertRedis(string(msg.Value))
	}
}

func insertRedis(body string){
	client := rejonson.ExtendClient(goRedisClient)
	client.JsonArrAppend("squid_game",".",body)
}

func main(){
	ctx := context.Background()
	consume(ctx)
}