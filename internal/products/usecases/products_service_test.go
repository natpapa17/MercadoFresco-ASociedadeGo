package usecases_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, float64, float64, float64, float64, int, float64, int, int, int) {
	return "valid_code", "valid_description", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1
}

func makeUpdateParams() (int, string, string, float64, float64, float64, float64, int, float64, int, int, int) {
	return 1, "update_code", "update_description", 2.0, 2.0, 2.0, 2.0, 2, 2.0, 2, 2, 2
}

func makeProduct() domain.Product {
	return domain.Product{
		Id:                               1,
		Product_Code:                     "valid_code",
		Description:                      "valid_description",
		Width:                            1.0,
		Height:                           1.0,
		Length:                           1.0,
		Net_Weight:                       1.0,
		Expiration_Rate:                  1,
		Recommended_Freezing_Temperature: 1.0,
		Freezing_Rate:                    1,
		Product_Type_Id:                  1,
		Seller_Id:                        1,
	}
}

func makeUpdateProduct() domain.Product {
	return domain.Product{
		Id:                               2,
		Product_Code:                     "update_code",
		Description:                      "update_description",
		Width:                            2.0,
		Height:                           2.0,
		Length:                           2.0,
		Net_Weight:                       2.0,
		Expiration_Rate:                  2,
		Recommended_Freezing_Temperature: 2.0,
		Freezing_Rate:                    2,
		Product_Type_Id:                  2,
		Seller_Id:                        2,
	}
}

func TestGetAll(t *testing.T) {
	mockProductRepository := mocks.NewRepositoryProduct(t)
	service := usecases.NewProductService(mockProductRepository)

	t.Run("Should call GetAll from Product Repository", func(t *testing.T) {
		mockProductRepository.On("GetAll").Return(domain.Products{makeProduct()}, nil).Once()

		service.GetAll()

		mockProductRepository.AssertCalled(t, "GetAll")
	})

	t.Run("Should return an error if GetAll from Product Repository returns an error", func(t *testing.T) {
		mockProductRepository.On("GetAll").Return(domain.Products{makeProduct()}, errors.New("any_error")).Once()

		_, err := service.GetAll()

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return a slice of products on success", func(t *testing.T) {
		mockProductRepository.
			On("GetAll").
			Return(domain.Products{makeProduct()}, nil).
			Once()

		ps, err := service.GetAll()

		assert.Equal(t, domain.Products{makeProduct()}, ps)
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	mockProductRepository := mocks.NewRepositoryProduct(t)
	service := usecases.NewProductService(mockProductRepository)

	t.Run("Should call GetById from Product Repository with correct ID", func(t *testing.T) {
		mockProductRepository.On("GetById", mock.AnythingOfType("int")).Return(makeProduct(), nil).Once()

		service.GetById(1)

		mockProductRepository.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return an error if GetById from Product Repository returns an error", func(t *testing.T) {
		mockProductRepository.On("GetById", mock.AnythingOfType("int")).Return(makeProduct(), errors.New("any_error")).Once()

		_, err := service.GetById(1)

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Product on success", func(t *testing.T) {
		mockProductRepository.On("GetById", mock.AnythingOfType("int")).Return(makeProduct(), nil).Once()

		p, err := service.GetById(1)

		assert.Equal(t, makeProduct(), p)
		assert.Nil(t, err)
	})
}

func TestCreate(t *testing.T) {
	mockProductRepository := mocks.NewRepositoryProduct(t)
	service := usecases.NewProductService(mockProductRepository)

	t.Run("Should Call GetByCode from Product Repository with correct code", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeProduct(), nil).Once()

		service.Create(makeCreateParams())

		mockProductRepository.AssertCalled(t, "GetByCode", "valid_code")
	})

	t.Run("Should return error if product code provided is in use", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(makeProduct(), nil).Once()

		_, err := service.Create(makeCreateParams())

		assert.EqualError(t, err, "product code: valid_code is already in use")
	})

	t.Run("Should call Create from Product Repository with correct values", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeProduct(), nil).Once()

		service.Create(makeCreateParams())

		mockProductRepository.AssertCalled(t, "Create", "valid_code", "valid_description", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
	})

	t.Run("Should return an error if Create from Product Repository returns an error", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(domain.Product{}, errors.New("any_error")).Once()

		_, err := service.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should create a Product with success", func(t *testing.T) {
		mockProductRepository.
			On("GetByCode", mock.AnythingOfType("string")).
			Return(domain.Product{}, errors.New("product code: valid_code is already in use")).
			Once()

		mockProductRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeProduct(), nil).
			Once()

		p, err := service.Create(makeCreateParams())

		assert.Equal(t, makeProduct(), p)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	mockProductRepository := mocks.NewRepositoryProduct(t)
	service := usecases.NewProductService(mockProductRepository)

	t.Run("Should call GetByCode from Product Repository with correct code", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeUpdateProduct(), nil).Once()

		service.Update(makeUpdateParams())

		mockProductRepository.AssertCalled(t, "GetByCode", "update_code")
	})

	t.Run("Should return conflict if product code is already in use", func(t *testing.T) {
		dbProduct := makeProduct()
		dbProduct.Id = 3
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(dbProduct, nil).Once()

		_, err := service.Update(makeUpdateParams())

		assert.EqualError(t, err, "product code: valid_code is already in use")
	})

	t.Run("Should call GetByCode from Product Repository with correct values", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeUpdateProduct(), nil).Once()

		service.Update(makeUpdateParams())

		mockProductRepository.AssertCalled(t, "GetByCode", "update_code")
	})

	t.Run("Should call Update from Product Repository with correct values", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeUpdateProduct(), nil).Once()

		service.Update(makeUpdateParams())

		mockProductRepository.AssertCalled(t, "Update", 1, "update_code", "update_description", 2.0, 2.0, 2.0, 2.0, 2, 2.0, 2, 2, 2)
	})

	t.Run("Should return error if Update from Product Repository returns an error", func(t *testing.T) {
		mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(domain.Product{}, errors.New("any_error")).Once()

		mockProductRepository.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(domain.Product{}, errors.New("any_error")).Once()

		_, err := service.Update(makeUpdateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Updated Product on success", func(t *testing.T) {
		mockProductRepository.
			On("GetByCode", mock.AnythingOfType("string")).
			Return(domain.Product{}, errors.New("product code in use")).
			Once()
		mockProductRepository.
			On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeUpdateProduct(), nil).
			Once()

		p, err := service.Update(makeUpdateParams())

		assert.Equal(t, makeUpdateProduct(), p)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	mockProductRepository := mocks.NewRepositoryProduct(t)
	service := usecases.NewProductService(mockProductRepository)

	t.Run("Should call Delete from Product Repository with correct ID", func(t *testing.T) {
		mockProductRepository.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

		service.Delete(1)

		mockProductRepository.AssertCalled(t, "Delete", 1)
	})

	t.Run("Should return an error if Delete from Product Repository returns an error", func(t *testing.T) {
		mockProductRepository.
			On("Delete", mock.AnythingOfType("int")).Return(errors.New("Error")).Once()

		err := service.Delete(1)

		assert.EqualError(t, err, "Error")
	})

	t.Run("Should Delete a Product with success", func(t *testing.T) {
		mockProductRepository.On("Delete", mock.AnythingOfType("int")).Return(nil)

		p := service.Delete(1)

		assert.Nil(t, p)
	})
}
