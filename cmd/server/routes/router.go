package routes

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses"
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

	warehouseFilePath, err := filepath.Abs("" + filepath.Join("data", "warehouses.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	warehouseFile := store.New(store.FileType, warehouseFilePath)
	wr := warehouses.CreateRepository(warehouseFile)
	ws := warehouses.CreateService(wr)
	wc := controllers.CreateWarehouseController(ws)

	sellersDb := store.New(store.FileType, "data/sellers.json")
	sellerRepo := sellers.NewRepository(sellersDb)
	sellerService := sellers.NewService(sellerRepo)
	sellerControllers := controllers.NewSeller(sellerService)

	sdb := store.New(store.FileType, "data/sections.json")
	sr := sections.NewRepository(sdb)
	ss := sections.NewService(sr)
	sc := section.NewSection(ss)

	mux := r.Group("api/")
	{
		warehouse := mux.Group("warehouses")
		{
			warehouse.GET("/", wc.GetAllWarehouses)
			warehouse.GET("/:id", wc.GetByIdWarehouse)
			warehouse.PATCH("/:id", wc.UpdateByIdWarehouse)
			warehouse.DELETE("/:id", wc.DeleteByIdWarehouse)
			warehouse.POST("/", wc.CreateWarehouse)
		}
		seller := mux.Group("seller")
		{
			seller.GET("/", sellerControllers.GetAll())
			seller.GET("/:id", sellerControllers.GetByIdSeller())
			seller.POST("/", sellerControllers.Store())
			seller.DELETE("/:id", sellerControllers.Delete())
			seller.PATCH("/:id", sellerControllers.Update())

		}

		sec := mux.Group("section")
		{
			sec.GET("/", sc.GetAll())
			sec.POST("/", sc.Add())
			sec.GET("/:id", sc.GetById())
			sec.PATCH("/:id", sc.UpdateById())
			sec.DELETE("/:id", sc.Delete())
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
