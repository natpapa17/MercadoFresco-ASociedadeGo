package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {
	mux := r.Group("api/")
	{
		warehouse := mux.Group("warehouse")
		{
			warehouse.GET("/", controllers.Ping)
		}
	}

	return r
}
