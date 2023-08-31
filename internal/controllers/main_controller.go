package controllers

import (
	"github.com/echpochmak31/avitotechbackendservice/internal/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

type MainController struct {
	address               string
	app                   *fiber.App
	synchronizationTicker *time.Ticker
	segmentService        services.AbstractSegmentService
	reportService         services.AbstractReportService
}

func InitMainController(service services.AbstractSegmentService,
	reportService services.AbstractReportService, address string) *MainController {
	mc := new(MainController)
	mc.app = fiber.New()
	mc.address = address
	mc.segmentService = service
	mc.reportService = reportService
	mc.synchronizationTicker = time.NewTicker(20 * time.Second)
	mc.setupRoutes()
	return mc
}

func (mc *MainController) Run() error {
	go mc.segmentService.SynchronizeSegments(mc.synchronizationTicker)
	return mc.app.Listen(mc.address)
}

func (mc *MainController) setupRoutes() {
	mc.app.Post("/segments", mc.createNewSegment)
	mc.app.Delete("/segments", mc.deleteSegment)
	mc.app.Get("/segments/user/:userId", mc.getActiveUserSegments)
	mc.app.Post("/segments/user", mc.setUserSegments)
	mc.app.Get("/reports/file/:reportName", mc.getReport)
	mc.app.Post("/reports/form", mc.formReport)
}
