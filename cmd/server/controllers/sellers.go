package controllers

import (
	"fmt"
	"net/http"
	
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/web"
)

type SellerController struct {
	service sellers.Service
}

func NewSeller(s sellers.Service) *SellerController {
	return &SellerController{
		service: s,
	}
}

func (c *SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		

		s, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
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
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}

		s, err := c.service.Store( req.Cid, req.CompanyName, req.Address, req.Telephone)
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
		fmt.Println(id)
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
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
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
			ctx.JSON(400, gin.H{"error": "O nome da empresa é obrigatório"})
			return
		}
		if req.Address== "" {
			ctx.JSON(400, gin.H{"error": "O endereco é obrigatório"})
			return
		}
		if req.Telephone == "" {
			ctx.JSON(400, gin.H{"error": "O telefone é obrigatório"})
			return
		}
		if req.Cid== 0 {
			ctx.JSON(400, gin.H{"error": "A cidade  é obrigatória"})
			return
		}

		s, err := c.service.Update(int(id), req.Cid, req.CompanyName, req.Address, req.Telephone)
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
	Cid int `json:"Cid" binding:"required"`
	CompanyName string `json:"CompanyName" binding:"required"`
	Address string `json:"Address" binding:"required"`
	Telephone string `json:"Telephone" binding:"required"`
}

