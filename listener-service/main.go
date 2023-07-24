package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main(){

	//try connect to rabbitMQ
	rabbitConn, err := connect()

	if err!=nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ")

	//start listening for message


}

func connect()(*amqp.Connection, error){
	var count float64
	var backOff = 1* time.Second
	var connection *amqp.Connection

	for{
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err!=nil{
			fmt.Println("RabbitMQ not yet ready....")
			count++
		}else{
			connection = c
			break
		}

		if count>5{
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(count),2))*time.Second
		log.Println("backing off...")
		time.Sleep(backOff)

		continue

	}

	return connection, nil
}