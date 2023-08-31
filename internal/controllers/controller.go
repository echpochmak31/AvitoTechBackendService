package controllers

import (
	"encoding/json"
	"github.com/echpochmak31/avitotechbackendservice/internal/controllers/requests"
	"github.com/echpochmak31/avitotechbackendservice/internal/controllers/responses"
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"github.com/echpochmak31/avitotechbackendservice/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (mc *MainController) getActiveUserSegments(c *fiber.Ctx) error {
	userId, err := strconv.ParseInt(c.Params("userId"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id path parameter")
	}

	segments, err := mc.segmentService.GetActiveUserSegments(userId)
	if err != nil {
		return err
	}
	response := responses.GetUserSegmentsResponse{SegmentSlugs: segments}
	return c.JSON(response)
}

func (mc *MainController) createNewSegment(c *fiber.Ctx) error {
	req := requests.CreateSegmentRequest{}
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err = mc.segmentService.CreateNewSegment(req.SegmentName, req.UserPercentage)
	if err != nil {
		return err
	}
	c.Status(fiber.StatusCreated)
	return nil
}

func (mc *MainController) deleteSegment(c *fiber.Ctx) error {
	segmentName := c.Get("slug", "")
	if segmentName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid segment name header")
	}
	return mc.segmentService.DeleteSegment(segmentName)
}

func (mc *MainController) setUserSegments(c *fiber.Ctx) error {
	req := requests.AddUserToSegmentsRequest{}
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	segmentsToAdd := make([]models.AbstractSegmentWithTTL, len(req.SegmentsToAdd))
	for i, requestSegment := range req.SegmentsToAdd {
		segmentsToAdd[i] = models.SimpleSegmentWithTTL{
			Name:           requestSegment.Name,
			ExpirationDate: requestSegment.ExpirationDate,
		}
	}
	return mc.segmentService.SetUserSegments(req.UserId, segmentsToAdd, req.SegmentsToRemove)
}

func (mc *MainController) getReport(c *fiber.Ctx) error {
	reportName := c.Params("reportName", "")
	if reportName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid report name parameter")
	}
	return mc.reportService.SendReport(services.NewFiberReportHandler(c, reportName))
}

func (mc *MainController) formReport(c *fiber.Ctx) error {
	req := requests.FormReportRequest{}
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	reportName, err := mc.reportService.FormReport(req.StartDate, req.EndDate)
	if err != nil {
		return err
	}
	response := responses.FormReportResponse{ReportUri: mc.address + "/reports/file/" + reportName}
	return c.JSON(response)
}
