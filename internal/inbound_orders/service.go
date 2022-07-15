package inbound_orders

type Inbound_orders_service_interface interface {
	Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (Inbound_orders, error)
	GetNumberOfOdersByEmployeeId(id int) (int, error)
}

type inbound_orders_service_struct struct {
	inboundRepository   Inbound_Orders_RepositoryInterface
	employeeRepository  EmployeeInboundInterface
	warehouseRepository WareHouseRepositoryInterface
}
