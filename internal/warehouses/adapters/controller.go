package adapters

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
)

type WarehouseController struct {
	service usecases.Service
}

func CreateWarehouseController(ws usecases.Service) *WarehouseController {
	return &WarehouseController{
		service: ws,
	}
}

func (wc *WarehouseController) CreateWarehouse(ctx *gin.Context) {
	var req warehouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	w, err := wc.service.Create(req.WarehouseCode, req.Address, req.Telephone, req.MinimumCapacity, req.MinimumTemperature)
	if err != nil {
		if isCustomError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
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
		if isCustomError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
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

	var req warehouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	w, err := wc.service.UpdateById(id, req.WarehouseCode, req.Address, req.Telephone, req.MinimumCapacity, req.MinimumTemperature)
	if err != nil {
		if isCustomError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": w,
	})
}

func (wc *WarehouseController) DeleteByIdWarehouse(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = wc.service.DeleteById(id)
	if err != nil {
		if isCustomError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

type warehouseRequest struct {
	WarehouseCode      string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required"`
}

func (wr *warehouseRequest) Validate() error {
	if strings.TrimSpace(wr.WarehouseCode) == "" {
		return errors.New("warehouse_code can't be empty")
	}

	if strings.TrimSpace(wr.Address) == "" {
		return errors.New("address can't be empty")

	}

	if strings.TrimSpace(wr.Telephone) == "" {
		return errors.New("telephone can't be empty")

	}

	if match, err := regexp.MatchString("^\\([1-9]{2}\\)\\s[0-9]{4,5}-[0-9]{4}$", wr.Telephone); err != nil || !match {
		return errors.New("telephone must respect the pattern (xx) xxxxx-xxxx or (xx) xxxx-xxxx")

	}

	if wr.MinimumCapacity <= 0 {
		return errors.New("minimum_capacity must be greater than 0")
	}

	return nil
}

func isCustomError(e error) bool {
	var be *usecases.BusinessRuleError
	var fe *usecases.NoElementFoundError

	if errors.As(e, &be) {
		return true
	}

	if errors.As(e, &fe) {
		return true
	}

	return false
}
