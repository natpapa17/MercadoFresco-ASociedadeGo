package routes

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/newLController"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/newController"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/buyer"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/product"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/record"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	carrier_factories "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/factories"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/products_rec"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/factories"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {

	pr := products.NewMysqlRepository(db.GetInstance())
	ps := products.NewProductService(pr)
	pc := product.NewProductController(ps)

	rr := records.NewMysqlRepository(db.GetInstance())
	rrp := products_rec.NewMysqlProductRepository(db.GetInstance())
	rs := records.NewRecordsService(rr, rrp)
	rc := record.NewRecordController(rs)

	BuyersFilePath, err := filepath.Abs("" + filepath.Join("data", "buyers.json"))
	if err != nil {
		log.Fatal("can't load buyers data file")
	}
	buyersFile := store.New(store.FileType, BuyersFilePath)
	br := buyers.CreateBuyerRepository(buyersFile)
	bs := buyers.CreateBuyerService(br)
	bc := buyer.CreateBuyerController(bs)

	warehouseController := factories.MakeWarehouseController()
	carrierController := carrier_factories.MakeCarrierController()



	sellerCont := newController.NewSellerController()

	sdb := store.New(store.FileType, "data/sections.json")
	sr := sections.NewRepository(sdb)
	ss := sections.NewService(sr)
	sc := section.NewSection(ss)

	localityController := newLController.NewLocalityController()

	
	employeeFilePath, err := filepath.Abs("" + filepath.Join("data", "employee.json"))
	if err != nil {
		log.Fatal("can't load employee data file")
	}
	employeeFile := store.New(store.FileType, employeeFilePath)
	er := employee.CreateRepository(employeeFile)
	warehouseFilePath, err := filepath.Abs("" + filepath.Join("data", "warehouses.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	warehouseFile := store.New(store.FileType, warehouseFilePath)
	wr := adapters.CreateWarehouseFileRepository(warehouseFile)
	es := employee.CreateService(er, wr)
	ec := controllers.CreateEmployeeController(es)

	mux := r.Group("api/v1")
	{
		warehouse := mux.Group("warehouses")
		{
			warehouse.GET("/", warehouseController.GetAllWarehouses)
			warehouse.GET("/:id", warehouseController.GetByIdWarehouse)
			warehouse.PATCH("/:id", warehouseController.UpdateByIdWarehouse)
			warehouse.DELETE("/:id", warehouseController.DeleteByIdWarehouse)
			warehouse.POST("/", warehouseController.CreateWarehouse)
		}

		buyer := mux.Group("buyers")
		{
			buyer.GET("/", bc.GetAllBuyers)
			buyer.GET("/:id", bc.GetBuyerById)
			buyer.PATCH("/:id", bc.UpdateBuyerById)
			buyer.DELETE("/:id", bc.DeleteBuyerById)
			buyer.POST("/", bc.CreateBuyer)
		}

		seller := mux.Group("seller")
		{
			seller.GET("/", sellerCont.GetAll())
			seller.GET("/:id", sellerCont.GetByIdSeller())
			seller.POST("/", sellerCont.Store())
			seller.DELETE("/:id", sellerCont.Delete())
			seller.PATCH("/:id", sellerCont.Update())


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

		locality := mux.Group("localities")
		{
			
			locality.GET("/", localityController.ReportAll())
			locality.GET("/:id", localityController.ReportById())
			locality.POST("/", localityController.Create())
		}
		employee := mux.Group("employees")
		{
			employee.GET("/", ec.GetAllEmployee)
			employee.GET("/:id", ec.GetByIdEmployee)
			employee.PATCH("/:id", ec.UpdateByIdEmployee)
			employee.DELETE("/:id", ec.DeleteByIdEmployee)
			employee.POST("/", ec.CreateEmployee)
		}

		carriers := mux.Group("carriers")
		{
			carriers.POST("/", carrierController.CreateCarrier)
			carriers.GET("/:id", carrierController.GetNumberOfCarriersPerLocality)
		}

		records := mux.Group("records")
		{
			records.GET("/", rc.GetRecordsPerProduct())
			records.POST("/:id", rc.Create())
		}

	}

	return r
}
