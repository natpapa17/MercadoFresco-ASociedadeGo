package usecases

import "errors"

var ErrWarehouseCodeInUse = errors.New("this warehouse_code is already in use")

var ErrNoElementFound = errors.New("can't find element")
