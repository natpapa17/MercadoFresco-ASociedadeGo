package section

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
)

type SectionRequest struct {
	SectionNumber       int     `json:"section_number" binding:"required"`
	CurrentTemperature  float32 `json:"current_temperature" binding:"required"`
	MinimumTemprarature float32 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity     int     `json:"current_capacity" binding:"required"`
	MinimumCapacity     int     `json:"minimum_capacity" binding:"required"`
	MaximumCapacity     int     `json:"maximum_capacity" binding:"required"`
	WarehouseID         int     `json:"warehouse_id" binding:"required"`
	ProductTypeID       int     `json:"product_type_id" binding:"required"`
}

type SectionController struct {
	service sections.Service
}

func NewSection(s sections.Service) *SectionController {
	return &SectionController{
		service: s,
	}
}

func (c SectionController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": s})
	}
}

func (c SectionController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		s, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}

func (c *SectionController) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var obj SectionRequest

		if err := ctx.ShouldBindJSON(&obj); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		s, err := c.service.Add(obj.SectionNumber, obj.CurrentTemperature, obj.MinimumTemprarature, obj.CurrentCapacity, obj.MinimumCapacity, obj.MaximumCapacity, obj.WarehouseID, obj.ProductTypeID)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}

func (c *SectionController) UpdateById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var obj sections.Section

		if err := ctx.ShouldBindJSON(&obj); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		s, err := c.service.UpdateById(id, obj)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, s)
	}
}

func (c *SectionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
