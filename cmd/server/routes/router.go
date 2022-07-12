package routes

import (
	product_batch2 "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/product_batch"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch/repository/mysql"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	sm "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections/repository/mysql"
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

	BuyersFilePath, err := filepath.Abs("" + filepath.Join("data", "buyers.json"))
	if err != nil {
		log.Fatal("can't load buyers data file")
	}
	buyersFile := store.New(store.FileType, BuyersFilePath)
	br := buyers.CreateBuyerRepository(buyersFile)
	bs := buyers.CreateBuyerService(br)
	bc := controllers.CreateBuyerController(bs)

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

	// Common
	mdb := db.GetInstance()
	// Section
	sr := sm.NewMySQLRepository(mdb)
	ss := sections.NewService(sr)
	sc := section.NewSection(ss)
	// Product Batch
	pbr := mysql.NewMySQLRepository(mdb)
	pbs := product_batch.NewService(pbr)
	pbc := product_batch2.NewSection(pbs)

	//Employee:
	employeeFilePath, err := filepath.Abs("" + filepath.Join("data", "employee.json"))
	if err != nil {
		log.Fatal("can't load employee data file")
	}
	employeeFile := store.New(store.FileType, employeeFilePath)
	er := employee.CreateRepository(employeeFile)
	es := employee.CreateService(er, wr)
	ec := controllers.CreateEmployeeController(es)

	mux := r.Group("api/v1")
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
			buyer.GET("/:id", bc.GetBuyer)
			buyer.PATCH("/:id", bc.UpdateBuyer)
			buyer.DELETE("/:id", bc.DeleteBuyer)
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
			sec.GET("/reportProducts", pbc.GetById())
		}

		pb := mux.Group("productBatches")
		{
			pb.POST("/", pbc.Add())
		}

		employee := mux.Group("employees")
		{
			employee.GET("/", ec.GetAllEmployee)
			employee.GET("/:id", ec.GetByIdEmployee)
			employee.PATCH("/:id", ec.UpdateByIdEmployee)
			employee.DELETE("/:id", ec.DeleteByIdEmployee)
			employee.POST("/", ec.CreateEmployee)
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
