package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardo-v-r/golang-to-do-list-api/domain/entity"
)

type TaskHandler struct {
	service entity.TaskService
}

func NewTaskHandler(service entity.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TaskHandler) Post(c *gin.Context) {
	var task entity.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = h.service.Post(&task)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, task)
}
