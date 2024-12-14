package main 

import (
	"fmt"
	"math/rand"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		//user-specific properties that you must set 
		"bootstrap.servers":"",
		"sasl.username": "",
		"sasl.password":""

		// fixed properties 
		"security.protocol":"sasl_ssl",
		
	})
}