package adapters

import (
	"errors"
	_ "fmt"
	"net/http"
	_ "regexp"
	_ "strconv"
	"strings"
	_ "strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"
)

type PurchaseOrderController struct {
	service usecases.PurchaseOrderService
}

func CreatePurchaseOrderController(pos usecases.PurchaseOrderService) *PurchaseOrderController {
	return &PurchaseOrderController{
		service: pos,
	}
}

func (poc *PurchaseOrderController) CreatePurchaseOrder(ctx *gin.Context) {
	var req purchaseOrdersRequest
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

	b, err := poc.service.Create(req.OrderNumber, req.OrderDate, req.TrackingCode, req.BuyerId, req.ProductRecordId, req.OrderStatusId)
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

// func (poc *PurchaseOrderController) GetPurchaseOrderById(ctx *gin.Context) {
// 	ids := ctx.QueryArray("id")

// 	for _, stringId := range ids {
// 		id, err := strconv.Atoi(stringId)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": "invalid id",
// 			})
// 			return
// 		}

// 		b, err := poc.service.GetPurchaseOrderById(id)
// 		if err != nil {
// 			if CustomError(err) {
// 				ctx.JSON(http.StatusNotFound, gin.H{
// 					"error": err.Error(),
// 				})
// 				return
// 			}
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "internal server error",
// 			})
// 			return
// 		}

// 		return

// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": po,
// 	})
// }

type purchaseOrdersRequest struct {
	OrderNumber     string `json:"order_number" binding:"required"`
	OrderDate       string `json:"order_date" binding:"required"`
	TrackingCode    string `json:"tracking_code" binding:"required"`
	BuyerId         int    `json:"buyer_id" binding:"required"`
	ProductRecordId int    `json:"product_record_id" binding:"required"`
	OrderStatusId   int    `json:"order_status_id" binding:"required"`
}

func (por *purchaseOrdersRequest) Validate() error {
	if strings.TrimSpace(por.OrderNumber) == "" {
		return errors.New("order number can't be empty")
	}

	if strings.TrimSpace(por.OrderDate) == "" {
		return errors.New("order date can't be empty")
	}

	if strings.TrimSpace(por.TrackingCode) == "" {
		return errors.New("tracking code can't be empty")
	}

	if por.BuyerId < 1 {
		return errors.New("buyer id can't be empty or smaller than 1")
	}

	if por.ProductRecordId < 1 {
		return errors.New("product record id can't be empty  or smaller than 1")
	}

	if por.OrderStatusId < 1 {
		return errors.New("order status can't be empty  or smaller than 1")
	}


	return nil
}

func CustomError(e error) bool {
	var be *usecases.BusinessRuleError
	var fe *usecases.NoElementInFileError

	if errors.As(e, &be) {
		return true
	}

	if errors.As(e, &fe) {
		return true
	}

	return false
}
