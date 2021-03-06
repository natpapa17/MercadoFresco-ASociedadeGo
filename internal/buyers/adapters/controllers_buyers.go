package adapters

import (
	"errors"
	"fmt"
	"net/http"
	_ "regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"
)

type BuyerController struct {
	service usecases.Service
}

func CreateBuyerController(bs usecases.Service) *BuyerController {
	return &BuyerController{
		service: bs,
	}
}

func (bc *BuyerController) CreateBuyer(ctx *gin.Context) {
	var req buyerRequest
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

	b, err := bc.service.Create(req.FirstName, req.LastName, req.Address, req.DocumentNumber)
	if err != nil {
		if CustomError(err) {
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
		"data": b,
	})
}

func (bc *BuyerController) GetAllBuyers(ctx *gin.Context) {
	b, err := bc.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": b,
	})
}

func (bc *BuyerController) GetBuyerById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	b, err := bc.service.GetBuyerById(id)
	if err != nil {
		if CustomError(err) {
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
		"data": b,
	})
}

func (bc *BuyerController) UpdateBuyerById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var req buyerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(err)
	if err := req.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := bc.service.UpdateBuyerById(id, req.FirstName, req.LastName, req.Address, req.DocumentNumber)
	if err != nil {
		if CustomError(err) {
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
		"data": b,
	})
}

func (bc *BuyerController) DeleteBuyerById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = bc.service.DeleteBuyerById(id)
	if err != nil {
		if CustomError(err) {
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

type buyerRequest struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	Address        string `json:"address" binding:"required"`
	DocumentNumber string `json:"document" binding:"required"`
}

func (br *buyerRequest) Validate() error {
	if strings.TrimSpace(br.FirstName) == "" {
		return errors.New("first name can't be empty")
	}

	if strings.TrimSpace(br.LastName) == "" {
		return errors.New("last name can't be empty")
	}

	if strings.TrimSpace(br.Address) == "" {
		return errors.New("address can't be empty")
	}

	if strings.TrimSpace(br.DocumentNumber) == "" {
		return errors.New("document number can't be empty")
	}

	return nil
}

func CustomError(e error) bool {
	var fe *usecases.ErrNoElementFound

	return errors.As(e, &fe)
// 	if errors.As(e, &fe) {
// 		return true
// 	}

// 	return false
}
