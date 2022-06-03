package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	wr := warehouses.CreateRepository()
	ws := warehouses.CreateService(wr)
	wc := controllers.CreateWarehouseController(ws)

	mux := r.Group("api/")
	{
		warehouse := mux.Group("warehouse")
		{
			warehouse.GET("/", controllers.Ping)
			warehouse.POST("/", wc.CreateWarehouse)
		}
	}

	return r
}
