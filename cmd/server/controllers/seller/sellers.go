package controllers

import (
	"fmt"
	"net/http"

	"strconv"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/services"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/web"
)

type SellerController struct {
	service services.Service
}

func NewSeller(s services.Service) *SellerController {
	return &SellerController{
		service: s,
	}
}

func (c *SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		

		s, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK,  s)
	}
}

func (c *SellerController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"message": "Invalid inputs. Please check your inputs"})
			return
		}

		s, err := c.service.Store( req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, s),
		)
	}
}

func (sl *SellerController) GetByIdSeller() gin.HandlerFunc{
	return func (ctx *gin.Context){
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}
	
		s, err := sl.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "can't find element with this id",
			})
			return
		}
	
		ctx.JSON(http.StatusOK, gin.H{
			"data": s,
		})
	}
	
}


func (c *SellerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))	
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.CompanyName== "" {
			ctx.JSON(400, gin.H{"error": "O nome da empresa ?? obrigat??rio"})
			return
		}
		if req.Address== "" {
			ctx.JSON(400, gin.H{"error": "O endereco ?? obrigat??rio"})
			return
		}
		if req.Telephone == "" {
			ctx.JSON(400, gin.H{"error": "O telefone ?? obrigat??rio"})
			return
		}
		if req.Cid== 0 {
			ctx.JSON(400, gin.H{"error": "O CID ?? obrigat??rio"})
			return
		}

		s, err := c.service.Update(int(id), req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityId)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, s)
	}
}



func (c *SellerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}

type request struct {
	Cid int `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	LocalityId int `json:"locality_id" binding:"required"`
}

