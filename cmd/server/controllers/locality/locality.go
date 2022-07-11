package locality

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/repository"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/services"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/web"
)
// sqlmock datadog

type LocalityController struct{
	service services.Service
}

func NewLocality(l services.Service) *LocalityController {
	return &LocalityController{
		service: l,
	}
}

func (l *LocalityController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req repository.LocalityRequestCreate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"message": "Invalid inputs. Please check your inputs"})
			return
		}

		s, err := l.service.Create( req.Name, req.Province_id, req.Country_id)
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


func (l *LocalityController) ReportAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		report, err := l.service.ReportAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report))
	}

}


func (l *LocalityController) ReportById() gin.HandlerFunc{
	return func(ctx *gin.Context){
		id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}
	report, err := l.service.ReportById(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report))
	}
}

