package controllers_test

import (
"bytes"
"errors"
"net/http"
"net/http/httptest"
"testing"

"github.com/gin-gonic/gin"
"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/mocks"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
)

func makeUnprocessableCreateAndUpdateBody() *bytes.Buffer {
return bytes.NewBuffer([]byte(`
        {
            "first_name": "John",
            "last_name": "Cena",
            "address": "Rua 2",
			"document_number": "123456"
        }
    `))
}

func makeValidCreateBody() *bytes.Buffer {
return bytes.NewBuffer([]byte(`
        {
			"first_name": "John",
            "last_name": "Cena",
            "address": "Rua 2",
			"document_number": "123456"
        }
    `))
}

func makeValidUpdateBody() *bytes.Buffer {
return bytes.NewBuffer([]byte(`
        {
            "first_name": "John",
            "last_name": "Cena",
            "address": "Rua 2",
			"document_number": "123456"
        }
    `))
}

type TestCase struct {
RequestBody      string
ExpectedResponse string
}

func makeInvalidBodiesTestCases() []TestCase {
return []TestCase{
{
RequestBody: `
            {
				"first_name": "",
				"last_name": "Cena",
				"address": "Rua 2",
				"document_number": "123456"
            }
            `,
ExpectedResponse: "{\"error\":\"first name can't be empty\"}",
},
{
RequestBody: `
            {
				"first_name": "John",
				"last_name": "",
				"address": "Rua 2",
				"document_number": "123456"
            }
            `,
ExpectedResponse: "{\"error\":\"last name can't be empty\"}",
}, {
RequestBody: `
            {
				"first_name": "John",
				"last_name": "Cena",
				"address": "",
				"document_number": "123456"            }
            `,
ExpectedResponse: "{\"error\":\"address can't be empty\"}",
}, {
RequestBody: `
            {
				"first_name": "John",
				"last_name": "Cena",
				"address": "Rua 2",
				"document_number": ""
            }
            `,
},
}

func makeBuyer() buyer.Buyer {
return buyer.Buyer{
ID:                               1,
FirstName:                     "valid_first_name",
LastName:                      "valid_last_name",
Address:                            "valid_address",
DocumentNumber:                     "valid_number",
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

func TestGetAllProduct(t *testing.T) {
gin.SetMode(gin.TestMode)

mockProductService := mocks.NewService(t)
controller := controllers.NewProductController(mockProductService)

r := gin.Default()
r.GET("/products", controller.GetAll())

t.Run("find_all_200", func(t *testing.T) {
mockProductService.On("GetAll").Return([]products.Product{makeProduct()}, nil).Once()
res := httptest.NewRecorder()
req, _ := http.NewRequest(http.MethodGet, "/products", nil)
r.ServeHTTP(res, req)

assert.Equal(t, http.StatusOK, res.Code)
assert.Equal(t, "[{\"id\":1,\"product_code\":\"valid_code\",\"description\":\"valid_description\",\"width\":1,\"height\":1,\"length\":1,\"net_weight\":1,\"expiration_rate\":1,\"recommended_freezing_temperature\":1,\"freezing_rate\":1,\"product_type_id\":1,\"seller_id\":1}]", res.Body.String())
})

}

func TestGetByIdProduct(t *testing.T) {
gin.SetMode(gin.TestMode)

mockProductService := mocks.NewService(t)
controller := controllers.NewProductController(mockProductService)

r := gin.Default()
r.GET("/products/:id", controller.GetById())

t.Run("find_by_id_existent_200", func(t *testing.T) {
mockProductService.On("GetById", mock.AnythingOfType("int")).Return(makeProduct(), nil).Once()
res := httptest.NewRecorder()

req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
r.ServeHTTP(res, req)

assert.Equal(t, http.StatusOK, res.Code)
assert.Equal(t, "{\"id\":1,\"product_code\":\"valid_code\",\"description\":\"valid_description\",\"width\":1,\"height\":1,\"length\":1,\"net_weight\":1,\"expiration_rate\":1,\"recommended_freezing_temperature\":1,\"freezing_rate\":1,\"product_type_id\":1,\"seller_id\":1}", res.Body.String())
})

t.Run("find_by_id_non_existent_ 404", func(t *testing.T) {
mockProductService.On("GetById", mock.AnythingOfType("int")).Return(makeProduct(), errors.New("Error")).Once()
res := httptest.NewRecorder()

req, _ := http.NewRequest(http.MethodGet, "/products/404", nil)
r.ServeHTTP(res, req)

assert.Equal(t, http.StatusNotFound, res.Code)
assert.Equal(t, "{\"error\":\"Error\"}", res.Body.String())
})
}

func TestCreateProduct(t *testing.T) {
gin.SetMode(gin.TestMode)

mockProductService := mocks.NewService(t)
controller := controllers.NewProductController(mockProductService)

r := gin.Default()
r.POST("/products", controller.Create())

t.Run("create_ok_201", func(t *testing.T) {
mockProductService.
On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
Return(makeProduct(), nil).Once()

res := httptest.NewRecorder()
req, _ := http.NewRequest(http.MethodPost, "/products", makeValidCreateBody())
r.ServeHTTP(res, req)

assert.Equal(t, http.StatusCreated, res.Code)
assert.Equal(t, "{\"id\":1,\"product_code\":\"valid_code\",\"description\":\"valid_description\",\"width\":1,\"height\":1,\"length\":1,\"net_weight\":1,\"expiration_rate\":1,\"recommended_freezing_temperature\":1,\"freezing_rate\":1,\"product_type_id\":1,\"seller_id\":1}", res.Body.String())
})

t.Run("creat_conflict_409", func(t *testing.T) {})

t.Run("create_fail_422", func(t *testing.T) {})
}

func TestUpdateProduct(t *testing.T) {
gin.SetMode(gin.TestMode)

mockProductService := mocks.NewService(t)
controller := controllers.NewProductController(mockProductService)

r := gin.Default()
r.PATCH("/products/:id", controller.Update())

t.Run("update_ok_200", func(t *testing.T) {

})

t.Run("update_non_existent_404", func(t *testing.T) {

})
}

func TestDeleteProduct(t *testing.T) {
gin.SetMode(gin.TestMode)

mockProductService := mocks.NewService(t)
controller := controllers.NewProductController(mockProductService)

r := gin.Default()
r.DELETE("/products/:id", controller.Delete())

t.Run("delete_ok_204", func(t *testing.T) {
mockProductService.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
res := httptest.NewRecorder()
req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)
r.ServeHTTP(res, req)

assert.Equal(t, http.StatusNoContent, res.Code)
assert.Empty(t, res.Body.String())
})

t.Run("delete_non_existent_404", func(t *testing.T) {
mockProductService.On("Delete", mock.AnythingOfType("int")).Return(errors.New("Error")).Once()
res := httptest.NewRecorder()
req, _ := http.NewRequest(http.MethodDelete, "/products/404", nil)
r.ServeHTTP(res, req)

assert.Equal(t, http.StatusNotFound, res.Code)
assert.Equal(t, "{\"error\":\"Error\"}", res.Body.String())
})
}