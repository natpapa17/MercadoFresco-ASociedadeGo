package adapters_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeUnprocessableCreateAndUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
		{
			"warehouse_code": "XPTO",
			"minimum_capacity": 10,
			"minimum_temperature": 8.7
		}
	`))
}

func makeValidCreateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
    "warehouse_code": "valid_code",
    "address": "valid_address",
    "telephone": "(44) 99909-9999",
    "minimum_capacity": 10,
    "minimum_temperature": 8.7
	}
`))
}

func makeValidUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
    "warehouse_code": "valid_code",
    "address": "updated_address",
    "telephone": "(44) 99909-9999",
    "minimum_capacity": 10,
    "minimum_temperature": 8.7
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
				"warehouse_code": "   ",
				"address": "Rua Brasil 870",
				"telephone": "(44) 9999-9999",
				"minimum_capacity": 10,
				"minimum_temperature": 8.7
			}
			`,
			ExpectedResponseBody: "{\"error\":\"warehouse_code can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"warehouse_code": "XPTO",
				"address": "    ",
				"telephone": "(44) 9999-9999",
				"minimum_capacity": 10,
				"minimum_temperature": 8.7
			}
			`,
			ExpectedResponseBody: "{\"error\":\"address can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"warehouse_code": "XPTO",
				"address": "Rua Brasil 870",
				"telephone": "  ",
				"minimum_capacity": 10,
				"minimum_temperature": 8.7
			}
			`,
			ExpectedResponseBody: "{\"error\":\"telephone can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"warehouse_code": "XPTO",
				"address": "Rua Brasil 870",
				"telephone": "999",
				"minimum_capacity": 10,
				"minimum_temperature": 8.7
			}
			`,
			ExpectedResponseBody: "{\"error\":\"telephone must respect the pattern (xx) xxxxx-xxxx or (xx) xxxx-xxxx\"}",
		},
		{
			RequestBody: `
			{
				"warehouse_code": "XPTO",
				"address": "Rua Brasil 870",
				"telephone": "(44) 9999-9999",
				"minimum_capacity": -10,
				"minimum_temperature": 8.7
			}
			`,
			ExpectedResponseBody: "{\"error\":\"minimum_capacity must be greater than 0\"}",
		},
	}
}

func makeDBWarehouse() domain.Warehouse {
	return domain.Warehouse{
		Id:                 1,
		WarehouseCode:      "valid_code",
		Address:            "valid_address",
		Telephone:          "(99) 99999-9999",
		MinimumCapacity:    10,
		MinimumTemperature: 5.0,
	}
}

func makeUpdatedDBWarehouse() domain.Warehouse {
	return domain.Warehouse{
		Id:                 1,
		WarehouseCode:      "valid_code",
		Address:            "updated_address",
		Telephone:          "(99) 99999-9999",
		MinimumCapacity:    10,
		MinimumTemperature: 5.0,
	}
}

func TestCreateWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockWarehouseService := mocks.NewService(t)
	sut := adapters.CreateWarehouseController(mockWarehouseService)

	r := gin.Default()
	r.POST("/warehouses", sut.CreateWarehouse)

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/warehouses", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call Create from Warehouse Service with correct values", func(t *testing.T) {
		mockWarehouseService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeDBWarehouse(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		mockWarehouseService.AssertCalled(t, "Create", "valid_code", "valid_address", "(44) 99909-9999", 10, 8.7)
	})

	t.Run("Should return an error and 400 status if Create from Warehouse Service returns an Business Rule error", func(t *testing.T) {
		mockWarehouseService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(domain.Warehouse{}, &usecases.BusinessRuleError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if Create from Warehouse Service did not returns an custom error", func(t *testing.T) {
		mockWarehouseService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(domain.Warehouse{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 201 status and data on success", func(t *testing.T) {
		mockWarehouseService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeDBWarehouse(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"warehouse_code\":\"valid_code\",\"address\":\"valid_address\",\"telephone\":\"(99) 99999-9999\",\"minimum_capacity\":10,\"minimum_temperature\":5}}", rr.Body.String())
	})
}

func TestGetAllWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockWarehouseService := mocks.NewService(t)
	sut := adapters.CreateWarehouseController(mockWarehouseService)

	r := gin.Default()
	r.GET("/warehouses", sut.GetAllWarehouses)

	t.Run("Should call GetAll from Warehouse Service", func(t *testing.T) {
		mockWarehouseService.On("GetAll").Return(domain.Warehouses{makeDBWarehouse()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses", nil)
		r.ServeHTTP(rr, req)

		mockWarehouseService.AssertNumberOfCalls(t, "GetAll", 1)
	})

	t.Run("Should return an error and 500 status if GetAll from Warehouse Service returns an error", func(t *testing.T) {
		mockWarehouseService.On("GetAll").Return(domain.Warehouses{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockWarehouseService.On("GetAll").Return(domain.Warehouses{makeDBWarehouse()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":[{\"id\":1,\"warehouse_code\":\"valid_code\",\"address\":\"valid_address\",\"telephone\":\"(99) 99999-9999\",\"minimum_capacity\":10,\"minimum_temperature\":5}]}", rr.Body.String())
	})
}

func TestGetByIdWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockWarehouseService := mocks.NewService(t)
	sut := adapters.CreateWarehouseController(mockWarehouseService)

	r := gin.Default()
	r.GET("/warehouses/:id", sut.GetByIdWarehouse)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses/invalid_id", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call GetById from Warehouse Service with correct id", func(t *testing.T) {
		mockWarehouseService.On("GetById", mock.AnythingOfType("int")).Return(makeDBWarehouse(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses/1", nil)
		r.ServeHTTP(rr, req)

		mockWarehouseService.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return an error and 404 status if GetById from Warehouse Service returns not find the correspondent element", func(t *testing.T) {
		mockWarehouseService.On("GetById", mock.AnythingOfType("int")).Return(domain.Warehouse{}, &usecases.NoElementFoundError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses/404", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if GetById from Warehouse Service returns an error", func(t *testing.T) {
		mockWarehouseService.On("GetById", mock.AnythingOfType("int")).Return(domain.Warehouse{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockWarehouseService.On("GetById", mock.AnythingOfType("int")).Return(makeDBWarehouse(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/warehouses/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"warehouse_code\":\"valid_code\",\"address\":\"valid_address\",\"telephone\":\"(99) 99999-9999\",\"minimum_capacity\":10,\"minimum_temperature\":5}}", rr.Body.String())
	})
}

func TestUpdateWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockWarehouseService := mocks.NewService(t)
	sut := adapters.CreateWarehouseController(mockWarehouseService)

	r := gin.Default()
	r.PATCH("/warehouses/:id", sut.UpdateByIdWarehouse)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/warehouses/invalid_id", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/warehouses/1", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, "/warehouses/1", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call UpdateById from Warehouse Service with correct values", func(t *testing.T) {
		mockWarehouseService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeUpdatedDBWarehouse(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/warehouses/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		mockWarehouseService.AssertCalled(t, "UpdateById", 1, "valid_code", "updated_address", "(44) 99909-9999", 10, 8.7)
	})

	t.Run("Should return an error and 404 status if UpdateById from Warehouse Service returns an Business Rule error", func(t *testing.T) {
		mockWarehouseService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(domain.Warehouse{}, &usecases.BusinessRuleError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/warehouses/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if UpdateById from Warehouse Service did not returns an custom error", func(t *testing.T) {
		mockWarehouseService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(domain.Warehouse{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/warehouses/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockWarehouseService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeUpdatedDBWarehouse(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/warehouses/1", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"warehouse_code\":\"valid_code\",\"address\":\"updated_address\",\"telephone\":\"(99) 99999-9999\",\"minimum_capacity\":10,\"minimum_temperature\":5}}", rr.Body.String())
	})
}

func TestDeleteByIdWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockWarehouseService := mocks.NewService(t)
	sut := adapters.CreateWarehouseController(mockWarehouseService)

	r := gin.Default()
	r.DELETE("/warehouses/:id", sut.DeleteByIdWarehouse)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/warehouses/invalid_id", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call DeleteById from Warehouse Service with correct id", func(t *testing.T) {
		mockWarehouseService.On("DeleteById", mock.AnythingOfType("int")).Return(nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/warehouses/1", nil)
		r.ServeHTTP(rr, req)

		mockWarehouseService.AssertCalled(t, "DeleteById", 1)
	})

	t.Run("Should return an error and 404 status if DeleteById from Warehouse Service returns not find the correspondent element", func(t *testing.T) {
		mockWarehouseService.On("DeleteById", mock.AnythingOfType("int")).Return(&usecases.NoElementFoundError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/warehouses/404", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if DeleteById from Warehouse Service returns an error", func(t *testing.T) {
		mockWarehouseService.On("DeleteById", mock.AnythingOfType("int")).Return(errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/warehouses/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockWarehouseService.On("DeleteById", mock.AnythingOfType("int")).Return(nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/warehouses/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}
