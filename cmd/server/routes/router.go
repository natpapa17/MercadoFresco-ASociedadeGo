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
	{
		buyers := mux.Group("buyers")
		{
			buyers.GET("/buyers", controllers.GetBuyers)
			buyers.GET("/buyers/:id", controllers.GetBuyers)
			buyers.POST("/buyers", controllers.GetBuyers)
			buyers.PATCH("/buyers/:id", controllers.GetBuyers)
			buyers.DELETE("/buyers/:id", controllers.GetBuyers)
		}
	}

	return r
}
