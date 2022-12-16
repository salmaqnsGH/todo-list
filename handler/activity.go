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
	var input activity.ActivityIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	activityDetail, err := h.service.GetActivityByID(input)
	if err != nil {
		errMessage := fmt.Sprintf("Activity with ID %v Not Found", input)

		response := helper.FormatNotFoundError(errMessage, activityDetail)
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
		response := helper.FormatBadRequest("title cannot be null", activity)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newActivity, err := h.service.CreateActivity(input)
	if err != nil {
		var activity activity.Activity
		response := helper.FormatBadRequest("title cannot be null", activity)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity.FormatCreateActivity(newActivity))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) DeleteActivity(c *gin.Context) {
	var input activity.ActivityIdInput
	activity := activity.Activity{}

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteActivity(input)
	if err != nil {
		errMessage := fmt.Sprintf("Activity with ID %v Not Found", input)

		response := helper.FormatNotFoundError(errMessage, activity)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity)
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) UpdateActivity(c *gin.Context) {
	var inputID activity.ActivityIdInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	activityById, err := h.service.GetActivityByID(inputID)
	if err != nil {
		errMessage := fmt.Sprintf("Activity with ID %v Not Found", inputID)

		response := helper.FormatNotFoundError(errMessage, activityById)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData activity.CreateActivityInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		fmt.Println("1", err)
		var activity activity.Activity
		response := helper.FormatBadRequest("title cannot be null", activity)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedActivity, err := h.service.UpdateActivity(inputID, inputData)
	if err != nil {
		errMessage := fmt.Sprintf("Activity with ID %v Not Found", inputID)

		response := helper.FormatNotFoundError(errMessage, updatedActivity)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", activity.FormatActivity(updatedActivity))
	c.JSON(http.StatusOK, response)
}
