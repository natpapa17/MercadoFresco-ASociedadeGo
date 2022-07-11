package adapters_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeUnprocessableCreateAndUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
		{
			"last_name": "last name",
			"address": "address",
			"document": "doc number"
		}
	`))
}

func makeValidCreateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
	"order_number": "123",
	"order_date": "01-01-2022",
	"tracking_code": "123",
	"buyer_id": 1,
	"product_record_id": 1,
	"order_status_id": 1
	}
`))
}

func makeValidUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
		"order_number": "123",
		"order_date": "01-01-2022",
		"tracking_code": "123",
		"buyer_id": 1,
		"product_record_id": 1,
		"order_status_id": 1
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
				"order_number": " ",
				"order_date": "01-01-2022",
				"tracking_code": "123",
				"buyer_id": 1,
				"product_record_id": 1,
				"order_status_id": 1,
			}
			`,
			ExpectedResponseBody: "{\"error\":\"order number can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"order_number": "123",
				"order_date": " ",
				"tracking_code": "123",
				"buyer_id": 1,
				"product_record_id": 1,
				"order_status_id": 1,
			}
			`,
			ExpectedResponseBody: "{\"error\":\"order date can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"order_number": "123",
				"order_date": "01-01-2022",
				"tracking_code": " ",
				"buyer_id": 1,
				"product_record_id": 1,
				"order_status_id": 1,
			}
			`,
			ExpectedResponseBody: "{\"error\":\"tracking code can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"order_number": "123",
				"order_date": "01-01-2022",
				"tracking_code": "123",
				"buyer_id": ,
				"product_record_id": 1,
				"order_status_id": 1,
			}
			`,
			ExpectedResponseBody: "{\"error\":\"buyer id can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"order_number": "123",
				"order_date": "01-01-2022",
				"tracking_code": "123",
				"buyer_id": 1,
				"product_record_id": ,
				"order_status_id": 1,
			}
			`,
			ExpectedResponseBody: "{\"error\":\"product record id can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"order_number": "123",
				"order_date": "01-01-2022",
				"tracking_code": "123",
				"buyer_id": 1,
				"product_record_id": 1,
				"order_status_id": ,
			}
			`,
			ExpectedResponseBody: "{\"error\":\"order status id can't be empty\"}",
		},
	}
}

func makeDBPurchaseOrder() purchaseOrders.PurchaseOrder {
	return purchaseOrders.PurchaseOrder{
		ID:             1,
		OrderNumber: "123",
		OrderDate: "01-01-2022",
		TrackingCode: "123",
		BuyerId: 1,
		ProductRecordId: 1,
		OrderStatusId: 1,
	}
}

func TestCreatePurchaseOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockPurchaseOrderService := mocks.NewService(t)
	sut := purchaseOrders.CreatePurchaseOrderController(mockPurchaseOrderService)

	r := gin.Default()
	r.POST("/purchaseOrders", sut.CreatePurchaseOrder)

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/purchaseOrders", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call Create from Purchase Orders Service with correct values", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(makeDBPurchaseOrder(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/purchaseOrders", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "Create", "order_number", "order_date", "tracking_code", "buyer_id", "product_record_id", "order_status_id")
	})

	t.Run("Should return an error and 500 status if Create from Purchase Orders Service did not returns an custom error", func(t *testing.T) {
		mockPurchaseOrdersService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(purchaseOrders.PurchaseOrder{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/purchaseOrders", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 201 status and data on success", func(t *testing.T) {
		mockPurchaseOrdersService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(makeDBPurchaseOrder(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/purchaseOrders", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1, \"order_number\":\"123\", \"order_date\":\"01-01-2022\", \"tracking_code\":\"123\", \"buyer_id\":1, \"product_record_id\":1, \"order_status_id\":1}}", rr.Body.String())
	})
}

