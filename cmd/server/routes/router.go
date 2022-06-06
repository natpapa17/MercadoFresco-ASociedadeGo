package routes

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {
	productsFilePath, err := filepath.Abs("" + filepath.Join("data", "products.json"))
	if err != nil {
		log.Fatal("can't load products data file")
	}
	productsFile := store.New(store.FileType, productsFilePath)
	pr := products.NewRepository(productsFile)
	ps := products.NewProductService(pr)
	pc := controllers.NewProductController(ps)
	mux := r.Group("api/")
	{
		warehouse := mux.Group("warehouse")
		{
			warehouse.GET("/", controllers.Ping)
		}
		products := mux.Group("products")
		{
			products.GET("/", pc.GetAll())
			products.GET("/:id", pc.GetById())
			products.POST("/", pc.Create())
			products.PATCH("/:id", pc.Update())
			products.DELETE("/:id", pc.Delete())
		}

	}

	return r
}
