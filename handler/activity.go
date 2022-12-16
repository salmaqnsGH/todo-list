package handler

import (
	"fmt"
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

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity.FormatActivities(activities))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) GetActivityById(c *gin.Context) {
	var input activity.GetActivityByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Not Found 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	activityDetail, err := h.service.GetActivityByID(input)
	if err != nil {
		errMessage := fmt.Sprintf("Activity with ID %v Not Found", input.ID)
		response := helper.FormatNotFoundError(errMessage, "Not Found", activityDetail)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity.FormatActivity(activityDetail))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) CreateActivity(c *gin.Context) {
	var input activity.CreateActivityInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var activity activity.Activity
		response := helper.FormatBadRequest(activity)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newActivity, err := h.service.CreateActivity(input)
	if err != nil {
		var activity activity.Activity
		response := helper.FormatBadRequest(activity)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity.FormatCreateActivity(newActivity))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) DeleteActivity(c *gin.Context) {
	var input activity.GetActivityByIdInput
	activity := activity.Activity{}

	err := c.ShouldBindUri(&input)
	err = h.service.DeleteActivity(input)
	if err != nil {
		errMessage := fmt.Sprintf("Activity with ID %v Not Found", input.ID)
		response := helper.FormatNotFoundError("Not Found", errMessage, activity)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity)
	c.JSON(http.StatusOK, response)
}
