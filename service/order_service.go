package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ayyasy123/Assignment_Day_26/model"
	"github.com/Ayyasy123/Assignment_Day_26/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrderHandler(c *gin.Context) {
	var orderDto model.OrderDto

	// Bind langsung ke model
	err := c.ShouldBind(&orderDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	order := orderDto.ToModel()
	err = repository.DB.Create(&order).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save order: %s", err.Error())),
		)
		return
	}

	orderDto.ID = order.ID
	orderDto.CreatedAt = order.CreatedAt
	orderDto.UpdatedAt = order.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", orderDto))
}

func ReadByIdOrderHandler(c *gin.Context) {
	// Menambil Id dari endpoint
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", idParam)),
		)
		return
	}

	var order model.Order
	err = repository.DB.First(&order, id).Error
	statusCodeError := http.StatusInternalServerError
	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCodeError = http.StatusNotFound
	}

	if err != nil {
		c.JSON(
			statusCodeError,
			model.NewFailedResponse(fmt.Sprintf("failed to get order: %s", err.Error())),
		)
		return
	}

	var orderDto model.OrderDto
	orderDto.FillFromModel(order)
	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", orderDto))
}
