package inbound_order

import (
	"errors"
	"fmt"
)

type Inbound_orders_service_interface interface {
	Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (Inbound_orders, error)
	GetNumberOfOdersByEmployeeId(employeeIds []int) (Inbound_orders_reports, error)
}

type inbound_orders_service_struct struct {
	inboundRepository      Inbound_Orders_RepositoryInterface
	employeeRepository     EmployeeInboundInterface
	warehouseRepository    WareHouseRepositoryInterface
	productBatchRepository ProductBatchRepositoryInterface
}

func CreateNewInboundService(ib Inbound_Orders_RepositoryInterface, er EmployeeInboundInterface, wr WareHouseRepositoryInterface, pr ProductBatchRepositoryInterface) Inbound_orders_service_interface {
	return &inbound_orders_service_struct{
		inboundRepository:      ib,
		employeeRepository:     er,
		warehouseRepository:    wr,
		productBatchRepository: pr,
	}
}

func (s *inbound_orders_service_struct) GetNumberOfOdersByEmployeeId(employeeIds []int) (Inbound_orders_reports, error) {
	reports := Inbound_orders_reports{}
	fmt.Println("oi")

	for _, employeeId := range employeeIds {
		employee, err := s.employeeRepository.GetById(employeeId)

		if errors.Is(err, ErrNoElementFound) {
			continue
		}

		if err != nil {
			return Inbound_orders_reports{}, err
		}

		ordersByEmployee, err := s.inboundRepository.GetNumberOfOdersByEmployeeId(employeeId)

		if err != nil {
			return Inbound_orders_reports{}, err
		}

		report := Inbound_orders_report{
			Id:                   employee.Id,
			Card_number_id:       employee.Card_number_id,
			First_name:           employee.First_name,
			Last_name:            employee.Last_name,
			Warehouse_id:         employee.Warehouse_id,
			Inbound_orders_count: ordersByEmployee,
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (s *inbound_orders_service_struct) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (Inbound_orders, error) {

	employeeExists, err := s.employeeRepository.CheckIfEmployeeExistById(employeeId)

	if err != nil {
		return Inbound_orders{}, err
	}

	if !employeeExists {
		return Inbound_orders{}, &NoElementInFileError{errors.New("no employee found with this id")}
	}

	productBatchIdExists, err := s.productBatchRepository.CheckProductBatchIfExistById(productBatchId)

	if err != nil {
		return Inbound_orders{}, err
	}

	if !productBatchIdExists {
		return Inbound_orders{}, &NoElementInFileError{errors.New("no productBatch found with this id")}
	}

	warehouseIdExists, err := s.warehouseRepository.CheckIfWareHouseExistById(warehouseId)

	if err != nil {
		return Inbound_orders{}, err
	}

	if !warehouseIdExists {
		return Inbound_orders{}, &NoElementInFileError{errors.New("no warehouse found with this id")}
	}

	inboundOrders, err := s.inboundRepository.Create(orderDate, orderNumber, employeeId, productBatchId, warehouseId)

	if err != nil {
		return Inbound_orders{}, err
	}

	return inboundOrders, nil

}
