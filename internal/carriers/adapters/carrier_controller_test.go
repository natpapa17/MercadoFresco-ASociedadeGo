package adapters_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCarrier(t *testing.T) {
	makeUnprocessableCreateBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"cid": "XPTO",
			"company_name": "valid_name",
			"address": "valid_address"
		}
	`))
	}

	type TestCase struct {
		RequestBody          string
		ExpectedResponseBody string
	}

	makeInvalidCreateBodiesTestCases := func() []TestCase {
		return []TestCase{
			{
				RequestBody: `
				{
					"cid": "  ",
					"company_name": "Company Name",
					"address": "Rua Brasil 870",
					"telephone": "(44) 9999-9999",
					"locality_id": 1
				}
				`,
				ExpectedResponseBody: "{\"error\":\"cid can't be empty\"}",
			},
			{
				RequestBody: `
				{
					"cid": "XPTO",
					"company_name": "    ",
					"address": "Rua Brasil 870",
					"telephone": "(44) 9999-9999",
					"locality_id": 1
				}
				`,
				ExpectedResponseBody: "{\"error\":\"company_name can't be empty\"}",
			},
			{
				RequestBody: `
				{
					"cid": "XPTO",
					"company_name": "Company Name",
					"address": "   ",
					"telephone": "(44) 9999-9999",
					"locality_id": 1
				}
				`,
				ExpectedResponseBody: "{\"error\":\"address can't be empty\"}",
			},
			{
				RequestBody: `
				{
					"cid": "XPTO",
					"company_name": "Company Name",
					"address": "Rua Brasil 870",
					"telephone": "  ",
					"locality_id": 1
				}
				`,
				ExpectedResponseBody: "{\"error\":\"telephone can't be empty\"}",
			},
			{
				RequestBody: `
				{
					"cid": "XPTO",
					"company_name": "Company Name",
					"address": "Rua Brasil 870",
					"telephone": "(44)99999999",
					"locality_id": 1
				}
				`,
				ExpectedResponseBody: "{\"error\":\"telephone must respect the pattern (xx) xxxxx-xxxx or (xx) xxxx-xxxx\"}",
			},
			{
				RequestBody: `
				{
					"cid": "XPTO",
					"company_name": "Company Name",
					"address": "Rua Brasil 870",
					"telephone": "(44) 9999-9999",
					"locality_id": -1
				}
				`,
				ExpectedResponseBody: "{\"error\":\"invalid locality_id\"}",
			},
		}
	}

	makeValidCreateBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"cid": "valid_cid",
			"company_name": "valid_company_name",
			"address": "valid_address",
			"telephone": "(44) 9999-9999",
			"locality_id": 1
		}
	`))
	}

	makeDbCarrier := func() domain.Carrier {
		return domain.Carrier{
			Id:          1,
			Cid:         "valid_cid",
			CompanyName: "valid_company_name",
			Address:     "valid_address",
			Telephone:   "(44) 9999-9999",
			LocalityId:  1,
		}
	}

	makeSut := func() (*gin.Engine, *mocks.CarrierService) {
		gin.SetMode(gin.TestMode)

		mockCarrierService := mocks.NewCarrierService(t)
		sut := adapters.CreateCarryController(mockCarrierService)

		r := gin.Default()
		r.POST("/carriers", sut.CreateCarrier)

		return r, mockCarrierService
	}

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		r, _ := makeSut()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/carriers", makeUnprocessableCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		r, _ := makeSut()
		testCases := makeInvalidCreateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/carriers", bytes.NewBuffer([]byte(tc.RequestBody)))
			r.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call Create from Carrier Service with correct values", func(t *testing.T) {
		r, mockCarrierService := makeSut()
		mockCarrierService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(makeDbCarrier(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/carriers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		mockCarrierService.AssertCalled(t, "Create", "valid_cid", "valid_company_name", "valid_address", "(44) 9999-9999", 1)
	})

	t.Run("Should return an error and 409 status if Carrier cid is in use", func(t *testing.T) {
		r, mockCarrierService := makeSut()
		mockCarrierService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(domain.Carrier{}, usecases.ErrCidInUse).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/carriers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		assert.Equal(t, "{\"error\":\"this cid is in use\"}", rr.Body.String())
	})

	t.Run("Should return an error and 400 status if locality id did not exists", func(t *testing.T) {
		r, mockCarrierService := makeSut()
		mockCarrierService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(domain.Carrier{}, usecases.ErrInvalidLocalityId).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/carriers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"this locality_id is invalid\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if Create from Warehouse Service did not returns an custom error", func(t *testing.T) {
		r, mockCarrierService := makeSut()
		mockCarrierService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(domain.Carrier{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/carriers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should 201 status and data on success", func(t *testing.T) {
		r, mockCarrierService := makeSut()
		mockCarrierService.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(makeDbCarrier(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/carriers", makeValidCreateBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"cid\":\"valid_cid\",\"company_name\":\"valid_company_name\",\"address\":\"valid_address\",\"telephone\":\"(44) 9999-9999\",\"locality_id\":1}}", rr.Body.String())
	})
}

func Test_GetNumberOfCarriersPerLocality(t *testing.T) {
	makeSut := func() (*gin.Engine, *mocks.CarrierService) {
		gin.SetMode(gin.TestMode)

		mockCarrierService := mocks.NewCarrierService(t)
		sut := adapters.CreateCarryController(mockCarrierService)

		server := gin.Default()
		server.GET("/carriers", sut.GetNumberOfCarriersPerLocality)

		return server, mockCarrierService
	}

	makeExpectedReportBodyResponse := func() string {
		return "{\"data\":[{\"locality_id\":1,\"locality_name\":\"any_name_1\",\"carriers_count\":1},{\"locality_id\":2,\"locality_name\":\"any_name_2\",\"carriers_count\":2},{\"locality_id\":3,\"locality_name\":\"any_name_3\",\"carriers_count\":3}]}"
	}

	makeReportsNumberOfCarriersPerLocality := func() domain.ReportsNumberOfCarriersPerLocality {
		return domain.ReportsNumberOfCarriersPerLocality{
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    1,
				LocalityName:  "any_name_1",
				CarriersCount: 1,
			},
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    2,
				LocalityName:  "any_name_2",
				CarriersCount: 2,
			},
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    3,
				LocalityName:  "any_name_3",
				CarriersCount: 3,
			},
		}
	}
	t.Run("Should call GetAllNumberOfCarriersPerLocality from Carrier Service if no id is provided", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetAllNumberOfCarriersPerLocality").Return(makeReportsNumberOfCarriersPerLocality, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers", nil)
		server.ServeHTTP(rr, req)

		mockCarrierService.AssertCalled(t, "GetAllNumberOfCarriersPerLocality")
		mockCarrierService.AssertNumberOfCalls(t, "GetAllNumberOfCarriersPerLocality", 1)
	})

	t.Run("Should return 500 status and error if GetAllNumberOfCarriersPerLocality from Carrier Service returns an error", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetAllNumberOfCarriersPerLocality").Return(domain.ReportsNumberOfCarriersPerLocality{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers", nil)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())

	})

	t.Run("Should return 200 status and reports on success", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetAllNumberOfCarriersPerLocality").Return(makeReportsNumberOfCarriersPerLocality, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers", nil)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, makeExpectedReportBodyResponse(), rr.Body.String())
	})

	t.Run("Should return 400 if invalid id is provided", func(t *testing.T) {
		server, _ := makeSut()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers?id=a", nil)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid locality_id\"}", rr.Body.String())
	})

	t.Run("Should call GetNumberOfCarriersPerLocalities with a slice of provided ids", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetNumberOfCarriersPerLocalities", mock.Anything).Return(makeReportsNumberOfCarriersPerLocality(), nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers?id=1&id=2&id=3", nil)
		server.ServeHTTP(rr, req)

		mockCarrierService.AssertNumberOfCalls(t, "GetNumberOfCarriersPerLocalities", 1)
		mockCarrierService.AssertCalled(t, "GetNumberOfCarriersPerLocalities", []int{1, 2, 3})
	})

	t.Run("Should return 500 status and error if GetNumberOfCarriersPerLocalities from Carrier Service returns an error", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetNumberOfCarriersPerLocalities", mock.Anything).Return(domain.ReportsNumberOfCarriersPerLocality{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers?id=1&id=2&id=3", nil)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", rr.Body.String())
	})

	t.Run("Should return 404 status and empty report if GetNumberOfCarriersPerLocalities from Carrier Service returns an empty slice", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetNumberOfCarriersPerLocalities", mock.Anything).Return(domain.ReportsNumberOfCarriersPerLocality{}, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers?id=1&id=2&id=3", nil)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"data\":[]}", rr.Body.String())
	})

	t.Run("Should return 200 status and reports on success", func(t *testing.T) {
		server, mockCarrierService := makeSut()
		mockCarrierService.On("GetNumberOfCarriersPerLocalities", mock.Anything).Return(makeReportsNumberOfCarriersPerLocality(), nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/carriers?id=1&id=2&id=3", nil)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, makeExpectedReportBodyResponse(), rr.Body.String())
	})
}
