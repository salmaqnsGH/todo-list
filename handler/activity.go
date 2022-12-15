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
