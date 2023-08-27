package main

import (
	"context"
	"github.com/echpochmak31/avitotechbackendservice/internal/configs"
	"github.com/echpochmak31/avitotechbackendservice/internal/controllers"
	"github.com/echpochmak31/avitotechbackendservice/internal/repositories"
	"github.com/echpochmak31/avitotechbackendservice/internal/services"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbURL, err := configs.GetPostgresUrl()
	if err != nil {
		log.Fatal(err)
	}

	rep, _ := repositories.NewPgxRepository(context.Background(), dbURL)
	defer rep.Close()

	service := services.NewSegmentService(rep)
	mc := controllers.InitMainController(service)

	err = mc.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Exit(0)
}
