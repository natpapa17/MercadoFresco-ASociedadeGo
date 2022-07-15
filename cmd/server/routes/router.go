package routes

import (
	"log"
	"path/filepath"

	controllers "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/inbound_orders"
	product_batch2 "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/product_batch"
	inboundOrdersService "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/inbound_order"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/newLController"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch/repository/mysql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/newController"

	"github.com/gin-gonic/gin"

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
	sm "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections/repository/mysql"
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
	//inbound_orders
	ioR := inboundOrdersService.Create_Inbound_Orders_MySQLRepository(db.GetInstance())
	eioR := inboundOrdersService.CreateEmployeeMysqlRepositoryInbound(db.GetInstance())
	wioR := inboundOrdersService.CreateWarehouseMySQLRepository(db.GetInstance())
	pbioR := inboundOrdersService.CreateNewMySQLRepositoryBatchProduct(db.GetInstance())
	ioS := inboundOrdersService.CreateNewInboundService(ioR, eioR, wioR, pbioR)
	ioC := inbound_orders.CreateNewInboundOrderController(ioS)

	localityController := newLController.NewLocalityController()

	wrwsql := employee.CreateWarehouseMySQLRepository(db.GetInstance())
	ersql := employee.CreateEmployeeMysqlRepository(db.GetInstance())
	es := employee.CreateService(ersql, wrwsql)
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
			sec.GET("/reportProducts", pbc.GetById())
		}

		pb := mux.Group("productBatches")
		{
			pb.POST("/", pbc.Add())
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
			locality.GET("/reportCarriers", carrierController.GetNumberOfCarriersPerLocality)
		}

		employee := mux.Group("employees")
		{
			employee.GET("/", ec.GetAllEmployee)
			employee.GET("/:id", ec.GetByIdEmployee)
			employee.PATCH("/:id", ec.UpdateByIdEmployee)
			employee.DELETE("/:id", ec.DeleteByIdEmployee)
			employee.POST("/", ec.CreateEmployee)
			employee.GET("/reportInboundOrders", ioC.GetNumberOfOdersByEmployeeId)
		}

		carriers := mux.Group("carriers")
		{
			carriers.POST("/", carrierController.CreateCarrier)
		}

		records := mux.Group("records")
		{
			records.GET("/", rc.GetRecordsPerProduct())
			records.POST("/:id", rc.Create())
		}

		inboundOrdersCreate := mux.Group("inboundOrders")
		{
			inboundOrdersCreate.POST("/", ioC.CreateInboundOrder)
		}
	}

	return r
}
