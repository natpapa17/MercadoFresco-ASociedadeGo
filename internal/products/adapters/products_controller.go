package adapters

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases"
)

type ProductController struct {
	service usecases.ServiceProduct
}

func NewProductController(p usecases.ServiceProduct) *ProductController {
	return &ProductController{
		service: p,
	}
}

func (c *ProductController) GetAllProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) GetByIdProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Create(req.Product_Code, req.Description, req.Width, req.Height, req.Length, req.Net_Weight, req.Expiration_Rate, req.Recommended_Freezing_Temperature, req.Freezing_Rate, req.Product_Type_Id, req.Seller_Id)

		if err != nil {
			if strings.Contains(err.Error(), "is already in use") {
				ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if err == nil {
			ctx.JSON(http.StatusCreated, p)
			return
		}
	}
}

func (c *ProductController) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid ID",
			})
			return
		}

		var req Request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Update(id, req.Product_Code, req.Description, req.Width, req.Height, req.Length, req.Net_Weight, req.Expiration_Rate, req.Recommended_Freezing_Temperature, req.Freezing_Rate, req.Product_Type_Id, req.Seller_Id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}

type Request struct {
	Product_Code                     string  `json:"product_code" binding:"required"`
	Description                      string  `json:"description" binding:"required"`
	Width                            float64 `json:"width" binding:"required"`
	Height                           float64 `json:"height" binding:"required"`
	Length                           float64 `json:"length" binding:"required"`
	Net_Weight                       float64 `json:"net_weight" binding:"required"`
	Expiration_Rate                  int     `json:"expiration_rate" binding:"required"`
	Recommended_Freezing_Temperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	Freezing_Rate                    int     `json:"freezing_rate" binding:"required"`
	Product_Type_Id                  int     `json:"product_type_id" binding:"required"`
	Seller_Id                        int     `json:"seller_id"`
}

func (pr *Request) Validate() error {
	if strings.TrimSpace(pr.Product_Code) == "" {
		return errors.New("product_code can't be empty")
	}

	if strings.TrimSpace(pr.Description) == "" {
		return errors.New("description can't be empty")
	}
	return nil
}
