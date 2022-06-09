package routes

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {
//Buyers:
	buyersFilePath, err := filepath.Abs("" + filepath.Join("data", "buyers.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	buyerFile := store.New(store.FileType, buyersFilePath)
	br := buyers.CreateRepository(buyerFile)
	bs := buyers.CreateService(br)
	bc := controllers.CreateBuyerController(bs)
  
  //WareHouses:
	warehouseFilePath, err := filepath.Abs("" + filepath.Join("data", "warehouses.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	warehouseFile := store.New(store.FileType, warehouseFilePath)
	wr := warehouses.CreateRepository(warehouseFile)
	ws := warehouses.CreateService(wr)
	wc := controllers.CreateWarehouseController(ws)

	//Employee:
	employeeFilePath, err := filepath.Abs("" + filepath.Join("data", "employee.json"))
	if err != nil {
		log.Fatal("can't load employee data file")
	}
	employeeFile := store.New(store.FileType, employeeFilePath)
	er := employee.CreateRepository(employeeFile)
	es := employee.CreateService(er, wr)
	ec := controllers.CreateEmployeeController(es)

//Sellers
  sellersDbPath, err := filepath.Abs("" + filepath.Join("data", "employee.json"))
	if err != nil {
		log.Fatal("can't load employee data file")
	}

	sellersDb := store.New(store.FileType, sellersDbPath)
	sellerRepo := sellers.NewRepository(sellersDb)
	sellerService := sellers.NewService(sellerRepo)
	sellerControllers := controllers.NewSeller(sellerService)

//Section
  sectionsPath, err := filepath.Abs("" + filepath.Join("data", "employee.json"))
	if err != nil {
		log.Fatal("can't load employee data file")
	}
	sdb := store.New(store.FileType, sectionsPath)
	sr := sections.NewRepository(sdb)
	ss := sections.NewService(sr)
	sc := section.NewSection(ss)

///////////////////////////////////////////////////////////////////////////////////////
	mux := r.Group("api/v1/")
	{
		warehouse := mux.Group("warehouses")
		{
			warehouse.GET("/", wc.GetAllWarehouses)
			warehouse.GET("/:id", wc.GetByIdWarehouse)
			warehouse.PATCH("/:id", wc.UpdateByIdWarehouse)
			warehouse.DELETE("/:id", wc.DeleteByIdWarehouse)
			warehouse.POST("/", wc.CreateWarehouse)
		}

		employee := mux.Group("employees")
		{
			employee.GET("/", ec.GetAllEmployee)
			employee.GET("/:id", ec.GetByIdEmployee)
			employee.PATCH("/:id", ec.UpdateByIdEmployee)
			employee.DELETE("/:id", ec.DeleteByIdEmployee)
			employee.POST("/", ec.CreateEmployee)

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