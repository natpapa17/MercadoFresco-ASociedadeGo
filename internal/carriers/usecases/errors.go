package usecases

import "errors"

var ErrCidInUse = errors.New("this cid is in use")

var ErrInvalidLocalityId = errors.New("this locality_id is invalid")

var ErrNoElementFound = errors.New("can't find element")
