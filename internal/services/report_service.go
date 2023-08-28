package services

import (
	"errors"
	"github.com/echpochmak31/avitotechbackendservice/internal/configs"
	"github.com/echpochmak31/avitotechbackendservice/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"path"
	"strconv"
	"time"
)

type ReportService struct {
	fsConfig         configs.FileSystemConfig
	reportRepository repositories.ReportsRepository
}

type FiberReportHandler struct {
	ctx        *fiber.Ctx
	reportName string
}

func NewFiberReportHandler(ctx *fiber.Ctx, reportName string) FiberReportHandler {
	f := FiberReportHandler{
		ctx:        ctx,
		reportName: reportName,
	}
	return f
}

func (f FiberReportHandler) Handle(any any) error {
	pathToReports, ok := any.(string)
	if ok {
		return f.ctx.SendFile(path.Join(pathToReports, f.reportName))
	}
	return errors.New("parameter type mismatch")
}

func NewReportService(rep repositories.ReportsRepository, fsConfig configs.FileSystemConfig) *ReportService {
	r := new(ReportService)
	r.reportRepository = rep
	r.fsConfig = fsConfig
	return r
}

func (r *ReportService) FormReport(startDate time.Time, endDate time.Time) (string, error) {
	reportName := r.GetReportName(startDate, endDate)
	reportPath := path.Join(r.fsConfig.VirtualPathToReports, reportName)
	err := r.reportRepository.MakeReportFile(startDate, endDate, reportPath)
	if err != nil {
		return "", err
	}
	return reportName, nil
}

func (r *ReportService) GetReportName(startDate time.Time, endDate time.Time) string {
	return strconv.FormatInt(startDate.Unix(), 10) + "_" + strconv.FormatInt(endDate.Unix(), 10) + ".csv"
}

func (r *ReportService) SendReport(handler AbstractReportHandler) error {
	return handler.Handle(r.fsConfig.PathToReports)
}
