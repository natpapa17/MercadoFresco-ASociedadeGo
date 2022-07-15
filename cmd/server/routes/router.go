package routes

import (
	"log"
	"path/filepath"

	EmployeeControllers "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/employee"
	product_batch2 "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/product_batch"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/newLController"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch/repository/mysql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/newController"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	purchase_adapter "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/adapters"
	purchase_usecases "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	carrier_factories "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/factories"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/product_factories"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/record_factories"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	sm "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections/repository/mysql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/factories"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {


	br := adapters.CreateBuyerMySQLRepository(db.GetInstance())
	bs := usecases.CreateBuyerService(br)
	bc := adapters.CreateBuyerController(bs)

	por := purchase_adapter.CreatePurchaseOrderMySQLRepository(db.GetInstance())
	pos := purchase_usecases.CreatePurchaseOrderService(por)
	poc := purchase_adapter.CreatePurchaseOrderController(pos)

	productsController := product_factories.MakeProductController()
	recordsController := record_factories.MakeRecordsController()

	warehouseController := factories.MakeWarehouseController()
	carrierController := carrier_factories.MakeCarrierController()

	sellerCont := newController.NewSellerController()

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
	wr := employee.CreateWarehouseRepository(warehouseFile)
	es := employee.CreateService(er, wr)
	ec := EmployeeControllers.CreateEmployeeController(es)

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
			sec.GET("/reportProducts", pbc.GetById())
		}

		pb := mux.Group("productBatches")
		{
			pb.POST("/", pbc.Add())
		}

		products := mux.Group("products")
		{
			products.GET("/", productsController.GetAllProduct())
			products.GET("/:id", productsController.GetByIdProduct())
			products.POST("/", productsController.CreateProduct())
			products.PATCH("/:id", productsController.UpdateProduct())
			products.DELETE("/:id", productsController.DeleteProduct())
		}

		po := mux.Group("purchaseOrders")
		{
			po.POST("/", poc.CreatePurchaseOrder)
		}
		locality := mux.Group("localities")
		{

			locality.GET("/", localityController.ReportAll())
			locality.GET("/:id", localityController.ReportById())
			locality.POST("/", localityController.Create())
			locality.GET("/reportCarriers", carrierController.GetNumberOfCarriersPerLocality)
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
		}

		records := mux.Group("records")
		{
			records.GET("/", recordsController.GetRecordsPerProduct())
			records.POST("/", recordsController.Create())
		}

	}

	return r
}
