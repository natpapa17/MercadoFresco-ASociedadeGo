package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeUnprocessableCreateBody() *bytes.Buffer {
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

func makeDBWarehouse() warehouses.Warehouse {
	return warehouses.Warehouse{
		Id:                 1,
		WarehouseCode:      "valid_code",
		Address:            "valid_address",
		Telephone:          "(99) 99999-9999",
		MinimumCapacity:    10,
		MinimumTemperature: 5.0,
	}
}

func TestCreateWarehouse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockWarehouseService := mocks.NewService(t)
	sut := controllers.CreateWarehouseController(mockWarehouseService)

	r := gin.Default()
	r.POST("/warehouses", sut.CreateWarehouse)

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeUnprocessableCreateBody())
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
		mockWarehouseService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(warehouses.Warehouse{}, &warehouses.BusinessRuleError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/warehouses", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if Create from Warehouse Service did not returns an custom error", func(t *testing.T) {
		mockWarehouseService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(warehouses.Warehouse{}, errors.New("any_message")).Once()
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
