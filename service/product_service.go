package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Ayyasy123/Assignment_Day_26/model"
	"github.com/Ayyasy123/Assignment_Day_26/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProductHandler(c *gin.Context) {
	var productDto model.ProductDTO

	// Bind langsung ke model
	err := c.ShouldBind(&productDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	product := productDto.ToModel()
	err = repository.DB.Create(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save product: %s", err.Error())),
		)
		return
	}

	productDto.ID = product.ID
	productDto.CreatedAt = product.CreatedAt
	productDto.UpdatedAt = product.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func ReadProductsHandler(c *gin.Context) {
	var products []model.Product

	query := `select * from products`
	filter := c.Query("filter")
	var args []any
	if filter != "" {
		query = fmt.Sprintf(
			"%s %s",
			query,
			"where name = ?",
		)
		args = append(args, filter)
	}

	err := repository.DB.Raw(query, args...).Scan(&products).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save product: %s", err.Error())),
		)
		return
	}
	var productsDtos []model.ProductDTO
	for _, product := range products {
		var productDto model.ProductDTO
		productDto.FillFromModel(product)
		productsDtos = append(productsDtos, productDto)
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productsDtos))
}

func ReadByIdProductHandler(c *gin.Context) {
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

	// mencari product by Id
	var product model.Product
	err = repository.DB.First(&product, id).Error
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

	var productDto model.ProductDTO
	productDto.FillFromModel(product)
	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func UpdateProductHandler(c *gin.Context) {
	var productDto model.ProductDTO
	err := c.ShouldBind(&productDto)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())),
		)
		return
	}

	// Menambil Id dari endpoint
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	// mencari product by Id
	var existingProduct model.Product
	err = repository.DB.First(&existingProduct, id).Error
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

	product := productDto.ToModel()
	product.ID = existingProduct.ID
	product.CreatedAt = existingProduct.CreatedAt
	product.UpdatedAt = existingProduct.UpdatedAt

	err = repository.DB.Save(&existingProduct).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())),
		)
		return
	}

	productDto.ID = existingProduct.ID
	productDto.CreatedAt = existingProduct.CreatedAt
	productDto.UpdatedAt = existingProduct.UpdatedAt

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func DeleteProductHandler(c *gin.Context) {
	// Menambil Id dari endpoint
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("invalid id: %s", err.Error())),
		)
		return
	}

	product := model.Product{ID: id}
	result := repository.DB.Delete(product)
	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to delete product: %s", result.Error.Error())),
		)
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", nil))
}

var productUploadDir = "uploads/products"

func UploadProductImageHandler(c *gin.Context) {
	formFile, file, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.NewFailedResponse(fmt.Sprintf("failed to get upload file: %s", err.Error())),
		)
		return
	}
	defer formFile.Close()

	name := c.PostForm("name")
	path := filepath.Join(productUploadDir, name)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			model.NewFailedResponse(fmt.Sprintf("failed to save file: %s", err.Error())),
		)
		return
	}

	var productDto model.ProductDTO
	productDto.ImagePath = &path
	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func UploadByIdProductImageHandler(c *gin.Context) {
	formFile, file, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewFailedResponse(fmt.Sprintf("failed to get upload file: %s", err.Error())))
		return
	}
	defer formFile.Close()

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, model.NewFailedResponse("name is required"))
		return
	}

	path := filepath.Join(productUploadDir, name)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewFailedResponse(fmt.Sprintf("failed to save file: %s", err.Error())))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewFailedResponse(fmt.Sprintf("invalid id: %s", c.Param("id"))))
		return
	}

	var existingProduct model.Product
	if err := repository.DB.First(&existingProduct, id).Error; err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, model.NewFailedResponse(fmt.Sprintf("failed to get product: %s", err.Error())))
		return
	}

	// Update only ImagePath
	existingProduct.ImagePath = model.ToNullString(&path)

	// Save updated product to database
	if err := repository.DB.Save(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())))
		return
	}

	// Prepare response DTO using FillFromModel
	var productDto model.ProductDTO
	productDto.FillFromModel(existingProduct)

	c.JSON(http.StatusOK, model.NewSuccessResponse("Success", productDto))
}

func DownloadProductImageHandler(c *gin.Context) {
	// Parse and validate product ID from URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewFailedResponse(fmt.Sprintf("invalid id: %s", c.Param("id"))))
		return
	}

	// Fetch existing product
	var product model.Product
	if err := repository.DB.First(&product, id).Error; err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, model.NewFailedResponse(fmt.Sprintf("failed to get product: %s", err.Error())))
		return
	}

	// Check if the product has a valid image path
	if !product.ImagePath.Valid {
		c.JSON(http.StatusNotFound, model.NewFailedResponse("product does not have an image"))
		return
	}

	imagePath := product.ImagePath.String

	// Check if the file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, model.NewFailedResponse("image file not found"))
		return
	}

	// Send the file to the postman
	c.File(imagePath)
}
