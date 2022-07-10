package adapters

import (
	"errors"
	_"fmt"
	"net/http"
	_ "regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"

)

type PurchaseOrderController struct {
	service usecases.Service
}

func CreatePurchaseOrderController(pos purchase_orders.Service) *PurchaseOrderController {
	return &BuyerController{
		service: pos,
	}
}

func (poc *PurchaseOrderController) CreatePurchaseOrder(ctx *gin.Context) {
	var req purchaseOrderRequest
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

	b, err := bc.service.Create(req.OrderNumber, req.OrderDate, req.TrackingCode, req.BuyerId, req.ProductRecordId, req.OrderStatusId)
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


type buyerRequest struct {
	ID int `json:"id"`
	OrderNumber string `json:"order_number"`
	OrderDate string `json:"order_date"`
	TrackingCode string `json:"tracking_code"`
	BuyerId int `json:"buyer_id"`
	ProductRecordId int `json:"product_record_id"`
	OrderStatusId int `json:"order_status_id"`
}

func (br *buyerRequest) Validate() error {
	if strings.TrimSpace(br.OrderNumber) == "" {
		return errors.New("order number can't be empty")
	}

	if strings.TrimSpace(br.OrderDate) == "" {
		return errors.New("order date can't be empty")
	}

	if strings.TrimSpace(br.TrackingCode) == "" {
		return errors.New("tracking code can't be empty")
	}

	if strings.TrimSpace(br.BuyerId) == "" {
		return errors.New("buyer id can't be empty")
	}

	if strings.TrimSpace(br.ProductRecordId) == "" {
		return errors.New("product record id can't be empty")
	}

	if strings.TrimSpace(br.OrderStatusId) == "" {
		return errors.New("order status can't be empty")
	}

	return nil
}

func CustomError(e error) bool {
	var be *buyers.BusinessRuleError
	var fe *buyers.NoElementInFileError

	if errors.As(e, &be) {
		return true
	}

	if errors.As(e, &fe) {
		return true
	}

	return false
}
