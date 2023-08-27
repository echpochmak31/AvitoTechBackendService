package controllers

import (
	"encoding/json"
	"github.com/echpochmak31/avitotechbackendservice/internal/controllers/requests"
	"github.com/echpochmak31/avitotechbackendservice/internal/controllers/responses"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (mc *MainController) getActiveUserSegments(c *fiber.Ctx) error {
	userId, err := strconv.ParseInt(c.Params("userId"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id path parameter")
	}

	segments, err := mc.service.GetActiveUserSegments(userId)
	if err != nil {
		return err
	}
	response := responses.GetUserSegmentsResponse{SegmentSlugs: segments}
	return c.JSON(response)
}

func (mc *MainController) createNewSegment(c *fiber.Ctx) error {
	segmentName := c.Get("slug", "")
	if segmentName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid segment name header")
	}
	err := mc.service.CreateNewSegment(segmentName)
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
	return mc.service.DeleteSegment(segmentName)
}

func (mc *MainController) setUserSegments(c *fiber.Ctx) error {
	req := requests.AddUserToSegmentsRequest{}
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	return mc.service.SetUserSegments(req.UserId, req.SegmentsToAdd, req.SegmentsToRemove)
}
