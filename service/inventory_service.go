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

func ReadByIdInventoryHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", idParam)),
		)
		return
	}

	var inventory model.Inventory
	err = repository.DB.First(&inventory, id).Error
	statusCodeError := http.StatusInternalServerError
	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCodeError = http.StatusNotFound
	}

	if err != nil {
		c.JSON(
			statusCodeError,
			model.NewFailedResponse(fmt.Sprintf("failed to get product: %s", err.Error())),
		)
		return
	}

	var inventoryDto model.InventoryDto
	inventoryDto.FillFromModel(inventory)
	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", inventory))
}

func UpdateInventoryHandler(c *gin.Context) {
	var inventoryDto model.InventoryDto
	err := c.ShouldBind(&inventoryDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	var existingInventory model.Inventory
	err = repository.DB.First(&existingInventory, id).Error
	statusCodeError := http.StatusInternalServerError
	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCodeError = http.StatusNotFound
	}

	if err != nil {
		c.JSON(
			statusCodeError,
			model.NewFailedResponse(fmt.Sprintf("failed to get product: %s", err.Error())),
		)
		return
	}

	inventory := inventoryDto.ToModel()
	inventory.ID = existingInventory.ID
	inventory.CreatedAt = existingInventory.CreatedAt
	inventory.UpdatedAt = existingInventory.UpdatedAt

	err = repository.DB.Save(&existingInventory).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())),
		)
		return
	}

	inventoryDto.ID = existingInventory.ID
	inventoryDto.CreatedAt = existingInventory.CreatedAt
	inventoryDto.UpdatedAt = existingInventory.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", inventoryDto))
}
