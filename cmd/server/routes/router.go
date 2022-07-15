package routes

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/product"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/warehouse"
	purchase_adapter "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/adapters"
	purchase_usecases "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"
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
	pc := product.NewProductController(ps)

	mdb := db.GetInstance()

	br := adapters.CreateBuyerMySQLRepository(mdb)
	bs := usecases.CreateBuyerService(br)
	bc := adapters.CreateBuyerController(bs)

	por := purchase_adapter.CreatePurchaseOrderMySQLRepository(mdb)
	pos := purchase_usecases.CreatePurchaseOrderService(por)
	poc := purchase_adapter.CreatePurchaseOrderController(pos)

	warehouseFilePath, err := filepath.Abs("" + filepath.Join("data", "warehouses.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	warehouseFile := store.New(store.FileType, warehouseFilePath)
	wr := warehouses.CreateRepository(warehouseFile)
	ws := warehouses.CreateService(wr)
	wc := warehouse.CreateWarehouseController(ws)

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

		buyer := mux.Group("buyers")
		{
			buyer.GET("/", bc.GetAllBuyers)
			buyer.GET("/:id", bc.GetBuyerById)
			// buyer.GET("/reportPurchaseOrders", poc.GetPurchaseOrderById)
			buyer.PATCH("/:id", bc.UpdateBuyerById)
			buyer.DELETE("/:id", bc.DeleteBuyerById)
			buyer.POST("/", bc.CreateBuyer)
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

		po := mux.Group("purchaseOrders")
		{
			po.POST("/", poc.CreatePurchaseOrder)
		}
	}

	return r
}
