package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeUnprocessableCreateAndUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
		{
			"first_name": "John",
			"last_name": "Cena",
			"address": "Rua 2"
		}
	`))
}

func makeValidCreateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
		"first_name": "John",
		"last_name": "Cena",
		"address": "Rua 2"
		"document_number": "1234623"
	}
`))
}

func makeValidUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
	"first_name": "John",
	"last_name": "Cena",
	"address": "Rua 7"
	"document_number": "1234623"
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
				"first_name": "",
				"last_name": "Cena",
				"address": "Rua 7"
				"document_number": "1234623"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"first name can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"first_name": "John",
				"last_name": "Cena",
				"address": ""
				"document_number": "1234623"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"address can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"first_name": "John",
				"last_name": "",
				"address": "Rua 7"
				"document_number": "1234623"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"last name can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"first_name": "John",
				"last_name": "Cena",
				"address": "Rua 7"
				"document_number": "1234623"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"document number can't be empty\"}",
		},
	}
}

func makeDBBuyer() buyers.Buyer {
	return buyers.Buyer{
		ID:             1,
		FirstName:      "first name",
		Address:        "address",
		LastName:       "last name",
		DocumentNumber: "2342430043",
	}
}

func makeUpdatedDBBuyer() buyers.Buyer {
	return buyers.Buyer{
		ID:             1,
		FirstName:      "first name",
		Address:        "address",
		LastName:       "last name",
		DocumentNumber: "2342430043",
	}
}

func TestCreateBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := controllers.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.POST("/buyers", sut.CreateBuyer)

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
			req, _ := http.NewRequest(http.MethodPost, "/buyers", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call Create from Warehouse Service with correct values", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "Create", "first name", "last name", "address", "2342430043")
	})

	t.Run("Should return an error and 400 status if Create from Warehouse Service returns an Business Rule error", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(buyers.Buyer{}, &buyers.BusinessRuleError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if Create from Warehouse Service did not returns an custom error", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(buyers.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 201 status and data on success", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"2342430043\"}}", rr.Body.String())
	})
}

func TestGetAllBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := controllers.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.GET("/buyers", sut.GetAllBuyers)

	t.Run("Should call GetAll from Buyers Service", func(t *testing.T) {
		mockBuyerService.On("GetAll").Return([]buyers.Buyer{makeDBBuyer()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers", nil)
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertNumberOfCalls(t, "GetAll", 1)
	})

	t.Run("Should return an error and 500 status if GetAll from Buyers Service returns an error", func(t *testing.T) {
		mockBuyerService.On("GetAll").Return([]buyers.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("GetAll").Return([]buyers.Buyer{makeDBBuyer()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"2342430043\"}}", rr.Body.String())
	})
}

func TestGetByIdBuyers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := controllers.CreateBuyersController(mockBuyerService)

	r := gin.Default()
	r.GET("/buyers/:id", sut.GetByIdBuyers)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/invalid_id", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call GetById from Buyer Service with correct id", func(t *testing.T) {
		mockBuyerService.On("GetById", mock.AnythingOfType("int")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return an error and 404 status if GetById from Warehouse Service returns not find the correspondent element", func(t *testing.T) {
		mockBuyerService.On("GetById", mock.AnythingOfType("int")).Return(buyers.Buyer{}, &buyers.NoElementInFileError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/404", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if GetById from Buyer Service returns an error", func(t *testing.T) {
		mockBuyerService.On("GetById", mock.AnythingOfType("int")).Return(buyers.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("GetById", mock.AnythingOfType("int")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"2342430043\"}}", rr.Body.String())
	})
}

func TestUpdateBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := controllers.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.PATCH("/buyers/:id", sut.UpdateByIdBuyer)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/invalid_id", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeUnprocessableCreateAndUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call UpdateById from Buyer Service with correct values", func(t *testing.T) {
		mockBuyerService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeUpdatedDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "UpdateById", 1, "first name", "last name", "address", "2342430043")
	})

	t.Run("Should return an error and 404 status if UpdateById from Buyer Service returns an Business Rule error", func(t *testing.T) {
		mockBuyerService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(buyers.Buyer{}, &buyers.BusinessRuleError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if UpdateById from Buyer Service did not returns an custom error", func(t *testing.T) {
		mockBuyerService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(buyers.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(makeUpdatedDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"2342430043\"}}", rr.Body.String())
	})
}

func TestDeleteByIdBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := controllers.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.DELETE("/buyers/:id", sut.DeleteByIdBuyer)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/invalid_id", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call DeleteById from Buyer Service with correct id", func(t *testing.T) {
		mockBuyerService.On("DeleteById", mock.AnythingOfType("int")).Return(nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "DeleteById", 1)
	})

	t.Run("Should return an error and 404 status if DeleteById from Buyer Service returns not find the correspondent element", func(t *testing.T) {
		mockBuyerService.On("DeleteById", mock.AnythingOfType("int")).Return(&buyers.NoElementInFileError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/404", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if DeleteById from Buyer Service returns an error", func(t *testing.T) {
		mockBuyerService.On("DeleteById", mock.AnythingOfType("int")).Return(errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("DeleteById", mock.AnythingOfType("int")).Return(nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}
