package record

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records"
)

type RecordController struct {
	service records.Service
}

func NewRecordController(p records.Service) *RecordController {
	return &RecordController{
		service: p,
	}
}

func (c *RecordController) GetRecordsPerProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := ctx.QueryArray("id")

		result := []records.ReportRecords{}

		for _, string_id := range ids {
			id, _ := strconv.Atoi(string_id)

			r, err := c.service.GetRecordsPerProduct(id)

			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}

			result = append(result, r)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (c *RecordController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if req.Last_Update_Date == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "A data de atualização é obrigatória"})
			return
		}
		if req.Purchase_Price == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "O valor da compra é obrigatório"})
			return
		}
		if req.Sale_Price == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "O valor de venda é obrigatório"})
			return
		}
		if req.Product_Id == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "O ID do produto é obrigatório"})
			return
		}

		r, err := c.service.Create(req.Last_Update_Date, req.Purchase_Price, req.Sale_Price, req.Product_Id)

		if err == nil {
			ctx.JSON(http.StatusCreated, gin.H{
				"data": r,
			})
			return
		}

		now := time.Now()
		currentDate := now.Format("2006-01-02")

		if req.Last_Update_Date < currentDate {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		}

		if err != nil {
			if strings.Contains(err.Error(), "is already in use") {
				ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, r)
	}
}

type Request struct {
	Id               int    `json:"id" binding:"required"`
	Last_Update_Date string `json:"last_update_date" binding:"required"`
	Purchase_Price   int    `json:"purchase_price" binding:"required"`
	Sale_Price       int    `json:"sale_price" binding:"required"`
	Product_Id       int    `json:"product_id" binding:"required"`
}
