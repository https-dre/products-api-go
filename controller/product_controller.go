package controller

import (
	"go-api/model"
	"go-api/usecase"
	"strconv"

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

func (p *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		ctx.JSON(400, model.Response{Message: "id cannot be null"})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, model.Response{Message: "product id should be an integer"})
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(500, model.Response{Message: "internal server error"})
		return
	}

	if product == nil {
		ctx.JSON(404, model.Response{Message: "product not found"})
		return
	}

	ctx.JSON(200, product)
}