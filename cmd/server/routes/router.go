package routes

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	buyersFilePath, err := filepath.Abs("" + filepath.Join("data", "buyers.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	buyerFile := store.New(store.FileType, buyersFilePath)
	br := buyers.CreateRepository(buyerFile)
	bs := buyers.CreateService(br)
	bc := controllers.CreateBuyerController(bs)

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
			buyers.GET("/", bc.GetAllBuyers)
			buyers.GET("/:id", bc.GetBuyer)
			buyers.POST("/", bc.SendBuyer)
			buyers.PATCH("/:id", bc.UpdateBuyer)
			buyers.DELETE("/:id", bc.DeleteBuyer)
		}
	}

	return r
}