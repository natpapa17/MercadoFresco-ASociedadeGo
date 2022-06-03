package controllers

import (
	"net/http"
	"strconv"

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

func (wc *WarehouseController) GetAllWarehouses(ctx *gin.Context) {
	w, err := wc.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": w,
	})
}

func (wc *WarehouseController) GetByIdWarehouse(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	w, err := wc.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "can't find element with this id",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": w,
	})
}

func (wc *WarehouseController) UpdateByIdWarehouse(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var req updateWarehouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid inputs",
		})
	}

	w, err := wc.service.UpdateById(id, req.WarehouseCode, req.Address, req.Telephone, req.MinimumCapacity, req.MinimumTemperature)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "can't find element with this id",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
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

type updateWarehouseRequest struct {
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required"`
}
