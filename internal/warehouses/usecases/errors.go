package usecases

type BusinessRuleError struct {
	Err error
}

func (b *BusinessRuleError) Error() string {
	return b.Err.Error()
}

type NoElementFoundError struct {
	Err error
}

func (b *NoElementFoundError) Error() string {
	return b.Err.Error()
}
