package main

import(
	"github.com/segmentio/kafka-go"
	"context"
	"fmt"
	"log"
	"os"
	"github.com/KromDaniel/rejonson"
	"github.com/go-redis/redis"
	"strings"
)

const(
	topic = "Game"
	brokerAddress = "localhost:9092"
)
var goRedisClient = redis.NewClient(&redis.Options{
	Addr: "",
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
	jsonString, err := client.JsonGet("INDEX").Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	body = strings.Replace(body, "{", "{\"request_number\":"+jsonString+",", 1)

	client.JsonArrAppend("squid_game",".",body)

	client.JsonNumIncrBy("INDEX", ".", 1)

}

func main(){
	ctx := context.Background()
	consume(ctx)
}