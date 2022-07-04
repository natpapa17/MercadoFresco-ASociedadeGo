package product_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	controllers "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/product"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/mocks"
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
			Return(makeProduct(), nil).
			Once()
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeValidCreateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "{\"id\":1,\"product_code\":\"valid_code\",\"description\":\"valid_description\",\"width\":1,\"height\":1,\"length\":1,\"net_weight\":1,\"expiration_rate\":1,\"recommended_freezing_temperature\":1,\"freezing_rate\":1,\"product_type_id\":1,\"seller_id\":1}", res.Body.String())
	})

	t.Run("creat_conflict_409", func(t *testing.T) {
		mockProductService.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("float64"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(products.Product{}, errors.New("product code: valid_code is already in use")).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("create_fail_422", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/products", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "{\"error\":")
	})

}

func TestUpdateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductService := mocks.NewService(t)
	controller := controllers.NewProductController(mockProductService)

	r := gin.Default()
	r.PATCH("/products/:id", controller.Update())

	t.Run("update_ok_200", func(t *testing.T) {
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

	t.Run("update_non_existent_404", func(t *testing.T) {
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
