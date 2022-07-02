package products_test

import (
"errors"
"testing"

"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/mocks"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, float64, float64, float64, float64, int, float64, int, int, int) {
return "valid_code", "valid_description", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1
}

func makeUpdateParams() (int, string, string, float64, float64, float64, float64, int, float64, int, int, int) {
return 1, "update_code", "update_description", 2.0, 2.0, 2.0, 2.0, 2, 2.0, 2, 2, 2
}

func makeProduct() products.Product {
return products.Product{
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

func makeUpdateProduct() products.Product {
return products.Product{
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
mockProductRepository := mocks.NewRepository(t)
service := products.NewProductService(mockProductRepository)

t.Run("find_all", func(t *testing.T) {
mockProductRepository.
On("GetAll").
Return([]products.Product{makeProduct()}, nil).
Once()

ps, err := service.GetAll()

assert.Equal(t, []products.Product{makeProduct()}, ps)
assert.Nil(t, err)
})
}

func TestGetById(t *testing.T) {
mockProductRepository := mocks.NewRepository(t)
service := products.NewProductService(mockProductRepository)

t.Run("find_by_id_non_existent", func(t *testing.T) {
mockProductRepository.On("GetById", mock.AnythingOfType("int")).
Return(makeProduct(), errors.New("Error")).Once()

_, err := service.GetById(1)

assert.EqualError(t, err, "Error")
})

t.Run("find_by_id_existent", func(t *testing.T) {
mockProductRepository.On("GetById", mock.AnythingOfType("int")).Return(makeProduct(), nil).Once()

p, err := service.GetById(1)

assert.Equal(t, makeProduct(), p)
assert.Nil(t, err)
})
}

func TestCreate(t *testing.T) {
mockProductRepository := mocks.NewRepository(t)
service := products.NewProductService(mockProductRepository)

t.Run("creat_conflict", func(t *testing.T) {
mockProductRepository.On("LastID").Return(1, nil).Once()
mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(makeProduct(), nil).Once()

_, err := service.Create(makeCreateParams())

assert.EqualError(t, err, "product code: valid_code is already in use")
})

t.Run("create_ok", func(t *testing.T) {
mockProductRepository.On("LastID").Return(1, nil).Once()
mockProductRepository.
On("GetByCode", mock.AnythingOfType("string")).
Return(products.Product{}, errors.New("product code: valid_code is already in use")).
Once()
mockProductRepository.
On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
Return(makeProduct(), nil).
Once()

p, err := service.Create(makeCreateParams())

assert.Equal(t, makeProduct(), p)
assert.Nil(t, err)
})
}

func TestUpdate(t *testing.T) {
mockProductRepository := mocks.NewRepository(t)
service := products.NewProductService(mockProductRepository)

t.Run("update_non_existent", func(t *testing.T) {
dbProduct := makeProduct()
dbProduct.Id = 3
mockProductRepository.On("GetByCode", mock.AnythingOfType("string")).Return(dbProduct, nil).Once()

_, err := service.Update(makeUpdateParams())

assert.EqualError(t, err, "product code: valid_code is already in use")

})

t.Run("update_ok", func(t *testing.T) {
mockProductRepository.
On("GetByCode", mock.AnythingOfType("string")).
Return(products.Product{}, errors.New("product code in use")).
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
mockProductRepository := mocks.NewRepository(t)
service := products.NewProductService(mockProductRepository)

t.Run("delete_non_existent", func(t *testing.T) {
mockProductRepository.
On("Delete", mock.AnythingOfType("int")).Return(errors.New("Error")).Once()

err := service.Delete(1)

assert.EqualError(t, err, "Error")
})

t.Run("delete_ok", func(t *testing.T) {
mockProductRepository.On("Delete", mock.AnythingOfType("int")).Return(nil)

p := service.Delete(1)

assert.Nil(t, p)
})
}