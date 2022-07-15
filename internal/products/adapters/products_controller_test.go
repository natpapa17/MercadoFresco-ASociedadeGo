package adapters_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeUnprocessableCreateAndUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
        {
            "product_code": "COD01",
            "description": "Água",
            "width": 1.5
        }
    `))
}

func makeValidCreateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
        {
            "product_code": "valid_code",
            "description": "valid_description",
            "width": 1.0,
            "height": 1.0,
            "length": 1.0,
            "net_weight": 1.0,
            "expiration_rate": 1,
            "recommended_freezing_temperature": 1.0,
            "freezing_rate": 1,
            "product_type_id": 1,
            "seller_id": 1
        }
    `))
}

func makeValidUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
        {
            "product_code": "update_code",
            "description": "update_description",
            "width": 2.0,
            "height": 2.0,
            "length": 2.0,
            "net_weight": 2.0,
            "expiration_rate": 2,
            "recommended_freezing_temperature": 2.0,
            "freezing_rate": 2,
            "product_type_id": 2,
            "seller_id": 2
        }
    `))
}

type TestCase struct {
	RequestBody          string
	ExpectedResponseBody string
}

func makeInvalidCreateAndUpdateBodiesTestCases() []TestCase {
	return []TestCase{
		{
			RequestBody: `
            {
                "product_code": " ",
                "description": "valid_description",
                "width": 1.0,
                "height": 1.0,
                "length": 1.0,
                "net_weight": 1.0,
                "expiration_rate": 1,
                "recommended_freezing_temperature": 1.0,
                "freezing_rate": 1,
                "product_type_id": 1,
                "seller_id": 1
            }
            `,
			ExpectedResponseBody: "{\"error\":\"product_code can't be empty\"}",
		},
		{
			RequestBody: `
            {
                "product_code": "valid_code",
                "description": " ",
                "width": 1.0,
                "height": 1.0,
                "length": 1.0,
                "net_weight": 1.0,
                "expiration_rate": 1,
                "recommended_freezing_temperature": 1.0,
                "freezing_rate": 1,
                "product_type_id": 1,
                "seller_id": 1
            }
            `,
			ExpectedResponseBody: "{\"error\":\"description can't be empty\"}",
		},
	}
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

func makeProducts() domain.Products {
	return domain.Products{
		makeProduct(),
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

func TestGetAllProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductService := mocks.NewServiceProduct(t)
	controller := adapters.NewProductController(mockProductService)

	r := gin.Default()
	r.GET("/products", controller.GetAllProduct())

	t.Run("Should call GetAll from Product Service", func(t *testing.T) {
		mockProductService.On("GetAll").Return(makeProducts(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		r.ServeHTTP(rr, req)

		mockProductService.AssertNumberOfCalls(t, "GetAll", 1)
	})

	t.Run("Should return an error and 500 status if GetAll from Product Service returns an error", func(t *testing.T) {
		mockProductService.On("GetAll").Return(domain.Products{}, errors.New("any_error")).Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", res.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockProductService.On("GetAll").Return(makeProducts, nil).Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "[{\"id\":1,\"product_code\":\"valid_code\",\"description\":\"valid_description\",\"width\":1,\"height\":1,\"length\":1,\"net_weight\":1,\"expiration_rate\":1,\"recommended_freezing_temperature\":1,\"freezing_rate\":1,\"product_type_id\":1,\"seller_id\":1}]", res.Body.String())
	})

}

func TestGetByIdProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductService := mocks.NewServiceProduct(t)
	controller := adapters.NewProductController(mockProductService)

	r := gin.Default()
	r.GET("/products/:id", controller.GetByIdProduct())

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

	mockProductService := mocks.NewServiceProduct(t)
	controller := adapters.NewProductController(mockProductService)

	r := gin.Default()
	r.POST("/products", controller.CreateProduct())

	t.Run("Should return an error and 422 status if body request contains unprocessable entity", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "{\"error\":")
	})

	t.Run("Should call Create from Product Service with correct values", func(t *testing.T) {
		mockProductService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeProduct(), nil).Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeValidCreateBody())
		r.ServeHTTP(res, req)

		mockProductService.AssertCalled(t, "Create", "valid_code", "valid_description", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			res := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(res, req)
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, tc.ExpectedResponseBody, res.Body.String())
		}
	})

	t.Run("Should return an error and 409 status if Product code is in use", func(t *testing.T) {
		mockProductService.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(domain.Product{}, errors.New("product code: valid_code is already in use")).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})

	t.Run("Should 201 status and create a Product on success", func(t *testing.T) {
		mockProductService.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeProduct(), nil).
			Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeValidCreateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "{\"id\":1,\"product_code\":\"valid_code\",\"description\":\"valid_description\",\"width\":1,\"height\":1,\"length\":1,\"net_weight\":1,\"expiration_rate\":1,\"recommended_freezing_temperature\":1,\"freezing_rate\":1,\"product_type_id\":1,\"seller_id\":1}", res.Body.String())
	})
}

func TestUpdateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductService := mocks.NewServiceProduct(t)
	controller := adapters.NewProductController(mockProductService)

	r := gin.Default()
	r.PATCH("/products/:id", controller.UpdateProduct())

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/products/invalid_id", makeValidUpdateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "{\"error\":\"invalid ID\"}", res.Body.String())
	})

	t.Run("Should return an error and 422 status if body request contains unprocessable entity", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/products/1", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 404 if can't find Product", func(t *testing.T) {
		mockProductService.
			On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeUpdateProduct(), errors.New("Error")).
			Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/products/2", makeValidUpdateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "{\"error\":\"Error\"}", res.Body.String())

	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			res := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(res, req)
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, tc.ExpectedResponseBody, res.Body.String())
		}
	})
	t.Run("Should update a Product on success", func(t *testing.T) {
		mockProductService.
			On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makeUpdateProduct(), nil).
			Once()

		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/products/2", makeValidUpdateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "{\"id\":2,\"product_code\":\"update_code\",\"description\":\"update_description\",\"width\":2,\"height\":2,\"length\":2,\"net_weight\":2,\"expiration_rate\":2,\"recommended_freezing_temperature\":2,\"freezing_rate\":2,\"product_type_id\":2,\"seller_id\":2}", res.Body.String())

	})

}

func TestDeleteProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductService := mocks.NewServiceProduct(t)
	controller := adapters.NewProductController(mockProductService)

	r := gin.Default()
	r.DELETE("/products/:id", controller.DeleteProduct())

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/products/invalid_id", nil)
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "{\"error\":\"invalid ID\"}", res.Body.String())
	})

	t.Run("Should delete a Product on success", func(t *testing.T) {
		mockProductService.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
		assert.Empty(t, res.Body.String())
	})

	t.Run("Should return an error and 404 status if Products is not found", func(t *testing.T) {
		mockProductService.On("Delete", mock.AnythingOfType("int")).Return(errors.New("Error")).Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/products/404", nil)
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "{\"error\":\"Error\"}", res.Body.String())
	})
}
