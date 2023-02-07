 package main
 import (
	"fmt"
	route "github.com/devAndre-Isaac/simulator-full-cycle/application/route"
)
import (
	"fmt"
	kafka "github.com/devAndre-Isaac/simulator-full-cycle/application/kafka"
	"github.com/devAndre-Isaac/simulator-full-cycle/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka.Produce(msg)
	}
}

