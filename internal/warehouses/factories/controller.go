package factories

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
)

func MakeWarehouseController() *adapters.WarehouseController {
	dataSource := "root:123@tcp(db)/fresh_market?parseTime=true"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("failed to connect to mysql")
	}
	wr := adapters.CreateMySQLRepository(conn)
	ws := usecases.CreateService(wr)
	wc := adapters.CreateWarehouseController(ws)

	return wc
}
