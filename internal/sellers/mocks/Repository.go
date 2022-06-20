package mocks

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/stretchr/testify/mock"
)

type SellerRepository struct {
	mock.Mock
}

func (s *SellerRepository) Delete(id int) error{
	args:= s.Called(id)
	var err error
	if rf, ok := args.Get(0).(func(int) error); ok{
		err = rf(id)
	}else{
		err = args.Error(0)
	}
	return err
}

func (s *SellerRepository)Update(id int,Cid int, CompanyName string, Addres, Telephone string) (sellers.Seller, error){
	args := s.Called(id, Cid, CompanyName, Addres, Telephone )
	var seller sellers.Seller

	if rf, ok := args.Get(0).(func(
		id, Cid int,
		CompanyName, Address, Telephone string,
		
	) sellers.Seller); ok {
		seller = rf(id, Cid, CompanyName, Addres, Telephone)
	} else {
		if args.Get(0) != nil {
			seller = args.Get(0).(sellers.Seller)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return seller, err

}

func (s *SellerRepository) GetAll() ([]sellers.Seller, error) {
	args := s.Called()

	var seller[]sellers.Seller

	if rf, ok := args.Get(0).(func() []sellers.Seller); ok {
		seller= rf()
	} else {
		if args.Get(0) != nil {
			seller= args.Get(0).([]sellers.Seller)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return seller, err
}

func (s *SellerRepository) GetById(id int) (sellers.Seller, error) {
	args := s.Called()

	var seller sellers.Seller
	if rf, ok := args.Get(0).(func(int) sellers.Seller); ok{
		seller = rf(id)
	}else{
		seller = args.Get(0).(sellers.Seller)
	}
	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return seller, err

}


func (s *SellerRepository) LastID() (int, error) {
	args := s.Called()

	var lastID int

	if rf, ok := args.Get(0).(func() int); ok {
		lastID = rf()
	} else {
		if args.Get(0) != nil {
			lastID = args.Get(0).(int)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return lastID, err
}


func (s *SellerRepository) Store(id , Cid int,CompanyName, Addres, Telephone string) (sellers.Seller, error) {
	args := s.Called()

	var seller sellers.Seller

	if rf, ok := args.Get(0).(func(id , Cid int, CompanyName, Addres, Telephone string) sellers.Seller); ok {
		seller = rf(id, Cid, CompanyName, Addres, Telephone)
	} else {
		if args.Get(0) != nil {
			seller = args.Get(0).(sellers.Seller)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return seller, err
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *SellerRepository {
	mock := &SellerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}