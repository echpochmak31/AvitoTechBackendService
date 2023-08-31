package application

import (
	"context"
	"github.com/echpochmak31/avitotechbackendservice/internal/configs"
	"github.com/echpochmak31/avitotechbackendservice/internal/controllers"
	"github.com/echpochmak31/avitotechbackendservice/internal/repositories"
	"github.com/echpochmak31/avitotechbackendservice/internal/services"
)

// RunSegmentsService Start listening and serving HTTP requests from the given address.
func RunSegmentsService(address string) error {
	dbURL, err := configs.GetPostgresUrl()
	if err != nil {
		return err
	}
	fsConfig, err := configs.GetFileSystemConfig()
	if err != nil {
		return err
	}

	repository, _ := repositories.NewPgxRepository(context.Background(), dbURL)
	defer repository.Close()

	segmentService := services.NewSegmentService(repository)
	reportService := services.NewReportService(repository, fsConfig)
	mc := controllers.InitMainController(segmentService, reportService, address)

	err = mc.Run()

	if err != nil {
		return err
	}

	return nil
}
