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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create activity", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newActivity, err := h.service.CreateActivity(input)
	if err != nil {
		response := helper.APIResponse("Failed to create activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create activity", http.StatusOK, "success", activity.FormatActivity(newActivity))
	c.JSON(http.StatusOK, response)
}

func (h *activityHandler) DeleteActivity(c *gin.Context) {
	var input activity.GetActivityByIdInput

	activity, err := h.service.GetActivityByID(input)
	if err != nil {
		response := helper.APIResponse("Failed get detail of activity", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(activity)
	fmt.Println(err)
	err = c.ShouldBindUri(&input)
	h.service.DeleteActivity(input)
	if err != nil {
		response := helper.APIResponse("Failed delete activity 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_deleted": true}
	response := helper.APIResponse("Successfully delete activity", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
