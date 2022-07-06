package adapters

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
)

type CarrierController struct {
	service usecases.CarrierService
}

func CreateCarryController(ws usecases.CarrierService) *CarrierController {
	return &CarrierController{
		service: ws,
	}
}

func (cc *CarrierController) CreateCarrier(ctx *gin.Context) {
	var req carryCreateRequest
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

	w, err := cc.service.Create(req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityId)

	if err == nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"data": w,
		})
		return
	}

	if errors.Is(err, usecases.ErrCidInUse) {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors.Is(err, usecases.ErrInvalidLocalityId) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
}

func (cc *CarrierController) GetNumberOfCarriersPerLocality(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid locality_id",
		})
		return
	}

	c, err := cc.service.GetNumberOfCarriersPerLocality(id)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": c,
		})
		return
	}

	if errors.Is(err, usecases.ErrInvalidLocalityId) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
}

type carryCreateRequest struct {
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int    `json:"locality_id" binding:"required"`
}

func (ccr *carryCreateRequest) Validate() error {
	if strings.TrimSpace(ccr.Cid) == "" {
		return errors.New("cid can't be empty")
	}

	if strings.TrimSpace(ccr.CompanyName) == "" {
		return errors.New("company_name can't be empty")
	}

	if strings.TrimSpace(ccr.Address) == "" {
		return errors.New("address can't be empty")
	}

	if strings.TrimSpace(ccr.Telephone) == "" {
		return errors.New("telephone can't be empty")
	}

	if match, err := regexp.MatchString("^\\([1-9]{2}\\)\\s[0-9]{4,5}-[0-9]{4}$", ccr.Telephone); err != nil || !match {
		return errors.New("telephone must respect the pattern (xx) xxxxx-xxxx or (xx) xxxx-xxxx")
	}

	if ccr.LocalityId <= 0 {
		return errors.New("invalid locality_id")
	}

	return nil
}
