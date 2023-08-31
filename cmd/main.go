package main

import (
	"github.com/echpochmak31/avitotechbackendservice/pkg/application"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	err := application.RunSegmentsService("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.Exit(0)
}
