package inbound_orders

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/inbound_order"
)

type Inbound_order_controller struct {
	service inbound_order.Inbound_orders_service_interface
}

type inboundCreateRequest struct {
	OrderDate      string `json:"order_date" binding:"required"`
	OrderNumber    string `json:"order_number" binding:"required"`
	EmployeeId     int    `json:"employee_id" binding:"required"`
	ProductBatchId int    `json:"product_batch_id" binding:"required"`
	WarehouseId    int    `json:"warehouse_id" binding:"required"`
}

func CreateNewInboundOrderController(ic inbound_order.Inbound_orders_service_interface) *Inbound_order_controller {
	return &Inbound_order_controller{
		service: ic,
	}
}

func (icr *inboundCreateRequest) Validate() error {
	if strings.TrimSpace(icr.OrderDate) == "" {
		return errors.New("order_date can't be empty")
	}

	if strings.TrimSpace(icr.OrderNumber) == "" {
		return errors.New("order_number can't be empty")
	}

	if icr.EmployeeId <= 0 {
		return errors.New("employee_id should be greater then 0")
	}

	if icr.ProductBatchId <= 0 {
		return errors.New("product_batch_id should be greater then 0")
	}

	if icr.WarehouseId <= 0 {
		return errors.New("warehouse_id should be greater then 0")
	}

	return nil
}

func (ic *Inbound_order_controller) CreateInboundOrder(ctx *gin.Context) {
	var req inboundCreateRequest
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

	io, err := ic.service.Create(req.OrderDate, req.OrderNumber, req.EmployeeId, req.ProductBatchId, req.WarehouseId)
	fmt.Println(err)
	if err == nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"data": io,
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": err.Error(),
	})
}

func (ic *Inbound_order_controller) GetNumberOfOdersByEmployeeId(ctx *gin.Context) {
	stringIds := ctx.QueryArray("id")
	ids := []int{}

	for _, stringId := range stringIds {
		id, err := strconv.Atoi(stringId)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "invalid id(s) call",
			})
			return
		}

		ids = append(ids, id)
	}

	reports, err := ic.service.GetNumberOfOdersByEmployeeId(ids)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(reports) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data": reports,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": reports,
	})

}
