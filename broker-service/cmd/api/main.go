package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Starting server at port ",webPort)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
