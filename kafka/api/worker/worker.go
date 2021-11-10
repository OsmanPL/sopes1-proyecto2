package main

import(
	"github.com/segmentio/kafka-go"
	"context"
	"fmt"
	"log"
	"os"
)

const(
	topic = "Game"
	brokerAddress = "localhost:9092"
)

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
	}
}

func main(){
	ctx := context.Background()
	consume(ctx)
}