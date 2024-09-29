package controller

import (
	"go-api/usecase"
	"go-api/model"
	"github.com/gin-gonic/gin"
)
type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (p * ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(200, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(400, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, insertedProduct)
}