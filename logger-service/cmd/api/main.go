package main

import (
	"context"
	"log"
	"logger-service/data"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct{
	Models data.Model
}




func main() {
	//connect to mongo
	mongoClient, err := connectToMongo()

	if err!=nil{
		log.Panic(err)
	}

	client = mongoClient

	ctx, cancel :=context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()


	defer func ()  {
		if err = client.Disconnect(ctx); err!=nil{
			log.Panic(err)
		}
	}()


	app := Config{
		Models: data.New(client),
	}

	//grpc

	app.gRPCListen()

	//go app.serve()

	// srv := &http.Server{
	// 	Addr: fmt.Sprintf(":%s",webPort),
	// 	Handler: app.routes(),
	// }

	// err =srv.ListenAndServe()

	// if err!=nil{
	// 	log.Panic(err)
	// }

}


// func (app *Config) serve(){
// 	srv := &http.Server{
// 		Addr: fmt.Sprintf(":%s",webPort),
// 		Handler: app.routes(),
// 	}

// 	err :=srv.ListenAndServe()

// 	if err!=nil{
// 		log.Panic(err)
// 	}
// }

func connectToMongo() (*mongo.Client, error){
	//create connection status
	clientoptions := options.Client().ApplyURI(mongoURL)

	clientoptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})


	//conncet
	c, err := mongo.Connect(context.TODO(), clientoptions)

	if err !=nil{
		log.Println("Error connectiong", err)
		return nil, err
	}

	return c, nil

}
