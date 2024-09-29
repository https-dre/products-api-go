package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controller"
	"go-api/usecase"
	"go-api/db"
	"go-api/repository"
)
func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}