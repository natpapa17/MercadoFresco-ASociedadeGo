package factories

import (
	"log"
	"path/filepath"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

func MakeWarehouseController() *adapters.WarehouseController {
	warehouseFilePath, err := filepath.Abs("" + filepath.Join("data", "warehouses.json"))
	if err != nil {
		log.Fatal("can't load warehouse data file")
	}
	warehouseFile := store.New(store.FileType, warehouseFilePath)
	wr := adapters.CreateFileRepository(warehouseFile)
	ws := usecases.CreateService(wr)
	wc := adapters.CreateWarehouseController(ws)

	return wc
}
