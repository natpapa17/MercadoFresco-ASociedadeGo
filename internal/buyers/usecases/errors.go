package usecases


type ErrNoElementFound struct {
	Err error
}

func (b *ErrNoElementFound) Error() string {
	return b.Err.Error()
}