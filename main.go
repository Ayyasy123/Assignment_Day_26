package main

import (
	"log"

	"github.com/Ayyasy123/Assignment_Day_26/repository"
	"github.com/Ayyasy123/Assignment_Day_26/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/assignment_day_26?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	repository.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// Routes Product
	r.POST("/product", service.CreateProductHandler)
	r.GET("/product", service.ReadProductsHandler)
	r.GET("/product/:id", service.ReadByIdProductHandler)
	r.PUT("/product/:id", service.UpdateProductHandler)
	r.DELETE("/product/:id", service.DeleteProductHandler)
	r.POST("/upload-product-image", service.UploadProductImageHandler)
	r.POST("/upload-product-image/:id", service.UploadByIdProductImageHandler)
	r.GET("/download-product-image/:id", service.DownloadProductImageHandler)

	// Routes inventory
	r.POST("/inventory", service.CreateProductHandler)
	r.GET("/inventory/:id", service.ReadByIdInventoryHandler)
	r.PUT("/inventory/:id", service.UpdateInventoryHandler)

	// Routes Order
	r.POST("/order", service.CreateOrderHandler)
	r.GET("/order/:id", service.ReadByIdOrderHandler)

	//Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
