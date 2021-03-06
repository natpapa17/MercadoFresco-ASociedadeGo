package adapters_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases/mocks"
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
		"first_name": "first name",
		"last_name": "last name",
		"address": "address",
		"document": "doc number"
	}
`))
}

func makeValidUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
	"first_name": "first name",
	"last_name": "last name",
	"address": "address",
	"document": "doc number"
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
				"first_name": " ",
				"last_name": "last name",
				"address": "address",
				"document": "doc number"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"first name can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"first_name": "first name",
				"last_name": "last name",
				"address": " ",
				"document": "doc number"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"address can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"first_name": "first name",
				"last_name": " ",
				"address": "address",
				"document": "doc number"
			}
			`,
			ExpectedResponseBody: "{\"error\":\"last name can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"first_name": "first name",
				"last_name": "last name",
				"address": "address",
				"document": " "
			}
			`,
			ExpectedResponseBody: "{\"error\":\"document number can't be empty\"}",
		},
	}
}

func makeDBBuyer() domain.Buyer {
	return domain.Buyer{
		ID:             1,
		FirstName:      "first name",
		Address:        "address",
		LastName:       "last name",
		DocumentNumber: "doc number",
	}
}

func makeUpdatedDBBuyer() domain.Buyer {
	return domain.Buyer{
		ID:             1,
		FirstName:      "first name",
		Address:        "address",
		LastName:       "last name",
		DocumentNumber: "doc number",
	}
}

func TestCreateBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := adapters.CreateBuyerController(mockBuyerService)

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

	t.Run("Should call Create from Buyers Service with correct values", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "Create", "first name", "last name", "address", "doc number")
	})

	t.Run("Should return an error and 500 status if Create from Buyers Service did not returns an custom error", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 201 status and data on success", func(t *testing.T) {
		mockBuyerService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/buyers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"doc number\"}}", rr.Body.String())
	})
}

func TestGetAllBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := adapters.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.GET("/buyers", sut.GetAllBuyers)

	t.Run("Should call GetAll from Buyers Service", func(t *testing.T) {
		mockBuyerService.On("GetAll").Return(domain.Buyers{makeDBBuyer()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers", nil)
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertNumberOfCalls(t, "GetAll", 1)
	})

	t.Run("Should return an error and 500 status if GetAll from Buyers Service returns an error", func(t *testing.T) {
		mockBuyerService.On("GetAll").Return(domain.Buyers{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("GetAll").Return(domain.Buyers{makeDBBuyer()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":[{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"doc number\"}]}", rr.Body.String())
	})
}

func TestGetBuyerById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := adapters.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.GET("/buyers/:id", sut.GetBuyerById)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/invalid_id", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call GetBuyerById from Buyer Service with correct id", func(t *testing.T) {
		mockBuyerService.On("GetBuyerById", mock.AnythingOfType("int")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "GetBuyerById", 1)
	})

	t.Run("Should return an error and 404 status if GetBuyerById from Buyers Service returns not find the correspondent element", func(t *testing.T) {
		mockBuyerService.On("GetBuyerById", mock.AnythingOfType("int")).Return(domain.Buyer{}, &usecases.ErrNoElementFound{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/404", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if GetBuyerById from Buyer Service returns an error", func(t *testing.T) {
		mockBuyerService.On("GetBuyerById", mock.AnythingOfType("int")).Return(domain.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("GetBuyerById", mock.AnythingOfType("int")).Return(makeDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"doc number\"}}", rr.Body.String())
	})
}

func TestUpdateBuyerById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := adapters.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.PATCH("/buyers/:id", sut.UpdateBuyerById)

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

	t.Run("Should call UpdateBuyerById from Buyer Service with correct values", func(t *testing.T) {
		mockBuyerService.On("UpdateBuyerById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(makeUpdatedDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "UpdateBuyerById", 1, "first name", "last name", "address", "doc number")
	})

	// t.Run("Should return an error and 404 status if UpdateBuyerById from Buyer Service returns an Business Rule error", func(t *testing.T) {
	// 	mockBuyerService.On("UpdateBuyerById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain.Buyers{}, usecases.).Once()
	// 	rr := httptest.NewRecorder()
	// 	req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidUpdateBody())
	// 	r.ServeHTTP(rr, req)

	// 	assert.Equal(t, http.StatusNotFound, rr.Code)
	// 	assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	// })

	t.Run("Should return an error and 500 status if UpdateBuyerById from Buyer Service did not returns an custom error", func(t *testing.T) {
		mockBuyerService.On("UpdateBuyerById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain.Buyer{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidUpdateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("UpdateBuyerById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(makeUpdatedDBBuyer(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/buyers/1", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"first_name\":\"first name\",\"last_name\":\"last name\",\"address\":\"address\",\"document_number\":\"doc number\"}}", rr.Body.String())
	})
}

func TestDeleteBuyerByIdBuyer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBuyerService := mocks.NewService(t)
	sut := adapters.CreateBuyerController(mockBuyerService)

	r := gin.Default()
	r.DELETE("/buyers/:id", sut.DeleteBuyerById)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/invalid_id", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call DeleteBuyerById from Buyer Service with correct id", func(t *testing.T) {
		mockBuyerService.On("DeleteBuyerById", mock.AnythingOfType("int")).Return(nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		mockBuyerService.AssertCalled(t, "DeleteBuyerById", 1)
	})

	t.Run("Should return an error and 404 status if DeleteBuyerById from Buyer Service returns not find the correspondent element", func(t *testing.T) {
		mockBuyerService.On("DeleteBuyerById", mock.AnythingOfType("int")).Return(&usecases.ErrNoElementFound{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/404", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if DeleteBuyerById from Buyer Service returns an error", func(t *testing.T) {
		mockBuyerService.On("DeleteBuyerById", mock.AnythingOfType("int")).Return(errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockBuyerService.On("DeleteBuyerById", mock.AnythingOfType("int")).Return(nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/buyers/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}
