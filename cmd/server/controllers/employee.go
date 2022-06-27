package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
)

type employeeRequest struct {
	Card_number_id int    `json:"card_number_id" binding:"required"`
	First_name     string `json:"first_name" binding:"required"`
	Last_name      string `json:"last_name" binding:"required"`
	Warehouse_id   int    `json:"warehouse_id" binding:"required"`
}

func (er *employeeRequest) Validate() error {

	if er.Card_number_id <= 0 {
		return errors.New("card_number_id must be greater than 0")
	}

	if strings.TrimSpace(er.First_name) == "" {
		return errors.New("first_name can't be empty")

	}

	if strings.TrimSpace(er.Last_name) == "" {
		return errors.New("last_name can't be empty")

	}

	if er.Warehouse_id <= 0 {
		return errors.New("wareHouse_id must be greater than 0")
	}

	return nil
}

type EmployeeController struct {
	service employee.EmployeeServiceInterface
}

func CreateEmployeeController(es employee.EmployeeServiceInterface) *EmployeeController {
	return &EmployeeController{
		service: es,
	}
}

func (ec *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var req employeeRequest
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

	e, err := ec.service.Create(req.Card_number_id, req.First_name, req.Last_name, req.Warehouse_id)
	if err != nil {
		if CustomError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": e,
	})
}

func (ec *EmployeeController) GetAllEmployee(ctx *gin.Context) {
	e, err := ec.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": e,
	})
}

func (ec *EmployeeController) GetByIdEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	e, err := ec.service.GetById(id)
	if err != nil {
		if CustomError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": e,
	})
}

func (ec *EmployeeController) UpdateByIdEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var req employeeRequest
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

	if _, err := ec.service.GetById(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	e, err := ec.service.UpdateById(id, req.Card_number_id, req.First_name, req.Last_name, req.Warehouse_id)
	if err != nil {
		if CustomError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": e,
	})
}

func (ec *EmployeeController) DeleteByIdEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = ec.service.DeleteById(id)
	if err != nil {
		if CustomError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
