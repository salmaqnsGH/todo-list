package handler

import (
	"net/http"
	"todo-list/activity"
	"todo-list/helper"

	"github.com/gin-gonic/gin"
)

type activityHandler struct {
	service activity.Service
}

func NewActivityHandler(service activity.Service) *activityHandler {
	return &activityHandler{service}
}

func (h *activityHandler) GetActivities(c *gin.Context) {
	activities, err := h.service.GetActivities()
	if err != nil {
		response := helper.APIResponse("Failed get activities", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Lists of activities", http.StatusOK, "success", activity.FormatActivities(activities))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) GetActivityById(c *gin.Context) {
	var input activity.GetActivityByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed get detail of activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	activityDetail, err := h.service.GetActivityByID(input)
	if err != nil {
		response := helper.APIResponse("Failed get detail of activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Lists of activities", http.StatusOK, "success", activity.FormatActivity(activityDetail))
	c.JSON(http.StatusOK, response)
}
