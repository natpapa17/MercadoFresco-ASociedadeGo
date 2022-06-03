package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses"
)

type WarehouseController struct {
	service warehouses.Service
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"value": "ok",
	})
}

func CreateWarehouseController(ws warehouses.Service) *WarehouseController {
	return &WarehouseController{
		service: ws,
	}
}

func (wc *WarehouseController) CreateWarehouse(ctx *gin.Context) {
	var req createWarehouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid inputs",
		})
	}

	w, err := wc.service.Create(req.WarehouseCode, req.Address, req.Telephone, req.MinimumCapacity, req.MinimumTemperature)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": w,
	})
}

type createWarehouseRequest struct {
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required"`
}
