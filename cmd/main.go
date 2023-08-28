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
	fsConfig, err := configs.GetFileSystemConfig()
	if err != nil {
		log.Fatal(err)
	}

	repository, _ := repositories.NewPgxRepository(context.Background(), dbURL)
	defer repository.Close()

	segmentService := services.NewSegmentService(repository)
	reportService := services.NewReportService(repository, fsConfig)
	mc := controllers.InitMainController(segmentService, reportService, "127.0.0.1:8080")

	err = mc.Run()

	if err != nil {
		log.Fatal(err)
	}
	defer os.Exit(0)
}
