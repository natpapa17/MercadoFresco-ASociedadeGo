package inbound_order

import (
	"errors"
)

type BusinessRuleError struct {
	Err error
}

func (b *BusinessRuleError) Error() string {
	return b.Err.Error()
}

type NoElementInFileError struct {
	Err error
}

func (b *NoElementInFileError) Error() string {
	return b.Err.Error()
}

var ErrNoElementFound = errors.New("can't find element")
