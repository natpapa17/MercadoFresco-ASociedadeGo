package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
)

type ProductController struct {
	service products.Service
}

func NewProductController(p products.Service) *ProductController {
	return &ProductController{
		service: p,
	}
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}
		ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, p))
	}
}

func (c *ProductController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Print("teste")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, p))
	}
}

func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if req.ProductCode == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O código do produto é obrigatório"})
			return
		}
		if req.Description == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A descrição do produto é obrigatória"})
			return
		}
		if req.Width == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A largura do produto é obrigatória"})
			return
		}
		if req.Length == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O comprimento do produto é obrigatória"})
			return
		}
		if req.NetWeight == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O peso do produto é obrigatório"})
			return
		}
		if req.ExpirationRate == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A data de vencimento do produto é obrigatória"})
			return
		}
		if req.RecommendedFreezingTemperature == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A temperatura de congelamento do produto é obrigatória"})
			return
		}
		if req.FreezingRate == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A taxa de congelamento do produto é obrigatória"})
			return
		}
		if req.ProductTypeId == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O tipo associado do produto é obrigatório"})
			return
		}

		p, err := c.service.Create(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			NewResponse(http.StatusOK, p),
		)
	}
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		p, err := c.service.Update(id, req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}

type Request struct {
	ProductCode                    string  `json:"product_code" binding:"required"`
	Description                    string  `json:"description" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"netweight" binding:"required"`
	ExpirationRate                 int     `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate                   int     `json:"freezing_rate" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id"`
}
