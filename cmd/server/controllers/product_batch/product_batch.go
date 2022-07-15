package product_batch

import (
	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	"net/http"
	"strconv"
)

type ProductBatchRequest struct {
	BatchNumber        int    `json:"batch_number" binding:"required"`
	CurrentQuantity    int    `json:"current_quantity" binding:"required"`
	CurrentTemperature int    `json:"current_temperature" binding:"required"`
	DueDate            string `json:"due_date" binding:"required"`
	InitialQuantity    int    `json:"initial_quantity" binding:"required"`
	ManufacturingDate  string `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  int    `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature int    `json:"minimum_temperature" binding:"required"`
	ProductID          int    `json:"product_id" binding:"required"`
	SectionID          int    `json:"section_id" binding:"required"`
}

type ProductBatchController struct {
	service product_batch.Service
}

func NewSection(s product_batch.Service) *ProductBatchController {
	return &ProductBatchController{
		service: s,
	}
}

func (c ProductBatchController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		prl, err := c.service.GetById(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": prl,
		})
	}
}

func (c *ProductBatchController) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pbr ProductBatchRequest

		if err := ctx.ShouldBindJSON(&pbr); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		s, err := c.service.Add(ctx, pbr.BatchNumber, pbr.CurrentQuantity, pbr.CurrentTemperature, pbr.DueDate, pbr.InitialQuantity, pbr.ManufacturingDate, pbr.ManufacturingHour, pbr.MinimumTemperature, pbr.ProductID, pbr.SectionID)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, s)
	}
}
