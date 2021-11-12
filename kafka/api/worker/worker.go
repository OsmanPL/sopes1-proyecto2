package main

import(
	"github.com/segmentio/kafka-go"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/KromDaniel/rejonson"
	"github.com/go-redis/redis"
)

const(
	topic = "Game"
	brokerAddress = "localhost:9092"
)
var (
	usr      = "admin"
	pwd      = "HQ03ilp&dOtj"
	host     = "34.132.83.52"
	port     = 27017
	database = "squid_game"
)
var goRedisClient = redis.NewClient(&redis.Options{
	Addr: "34.121.89.39:6379",
  })
  

  type Game struct {
	ID       int    `json:"id"`
	GameName string `json:"gameName"`
	Winner   string `json:"winner"`
	Players  int    `json:"players"`
	Worker   string `json:"worker"`
}


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
		go insertRedis(string(msg.Value))
		go insertMongo(ctx,string(msg.Value))
	}
}

func insertRedis(body string){
	client := rejonson.ExtendClient(goRedisClient)
	client.JsonArrAppend("squid_game",".",body)
}

func insertMongo(ctx context.Context,body string){
	var game Game
	json.Unmarshal([]byte(body), &game)
	var collection = GetCollection("games")
	var err error
	_, err = collection.InsertOne(ctx, game)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}


func main(){
	ctx := context.Background()
	consume(ctx)
}