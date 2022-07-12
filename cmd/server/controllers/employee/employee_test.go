package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	employeeController "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeValidCreateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
    "card_number_id": 568,
    "first_name": "Valid_Name",
    "last_name": "Valid_Last_Name",
    "warehouse_id": 1
	}
`))
}

func makeDBEmployee() employee.Employee {
	return employee.Employee{
		Id:             1,
		Card_number_id: 568,
		First_name:     "Valid_Name",
		Last_name:      "Valid_Last_Name",
		Warehouse_id:   1,
	}
}

func makeUnprocessableCreateAndUpdateBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
		"last_name": "Valid_Lastt_Name",
		"warehouse_id": 1
	}`))
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
						"card_number_id": 0,
						"first_name": "Marcos",
						"last_name": "Mantovani",
						"warehouse_id": 1
					}
				`,
			ExpectedResponseBody: "{\"error\":\"card_number_id must be greater than 0\"}",
		},
		{
			RequestBody: `
				{
					"card_number_id": 1,
					"first_name": " ",
					"last_name": "Mantovani",
					"warehouse_id": 1
				}
			`,
			ExpectedResponseBody: "{\"error\":\"first_name can't be empty\"}",
		},
		{
			RequestBody: `
			{
				"card_number_id": 1,
				"first_name": "Marcos",
				"last_name": " ",
				"warehouse_id": 1
			}
		`,
			ExpectedResponseBody: "{\"error\":\"last_name can't be empty\"}",
		},
		{
			RequestBody: `
		{
			"card_number_id": 1,
			"first_name": "Marcos",
			"last_name": "Mantovani",
			"warehouse_id": 0
		}
	`,
			ExpectedResponseBody: "{\"error\":\"wareHouse_id must be greater than 0\"}",
		},
	}
}

func TestCreateEmployee(t *testing.T) {

	mockEmployeeService := mocks.NewEmployeeServiceInterface(t)
	sut := employeeController.CreateEmployeeController(mockEmployeeService)
	response := gin.Default()

	response.POST("/employees", sut.CreateEmployee)

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/employees", makeUnprocessableCreateAndUpdateBody())
		response.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data or length", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer([]byte(tc.RequestBody)))
			response.ServeHTTP(rr, req)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should Create from Employee Service with correct Values", func(t *testing.T) {
		mockEmployeeService.On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(employee.Employee{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/employees", makeValidCreateBody())
		response.ServeHTTP(rr, req)
	})

	t.Run("Should return an error and 400 status if Create from Employee Service returns an Business Rule Error", func(t *testing.T) {
		mockEmployeeService.On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(employee.Employee{}, &employee.BusinessRuleError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/employees", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if Create from employee Service did not returns an custom error", func(t *testing.T) {
		mockEmployeeService.On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(employee.Employee{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/employees", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return 201 status and data on sucess at creating employee", func(t *testing.T) {
		mockEmployeeService.On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(makeDBEmployee(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/employees", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"card_number_id\":568,\"first_name\":\"Valid_Name\",\"Last_name\":\"Valid_Last_Name\",\"warehouse_id\":1}}", rr.Body.String())
	})
}

func TestUpdateEmployee(t *testing.T) {
	mockEmployeeService := mocks.NewEmployeeServiceInterface(t)
	sut := employeeController.CreateEmployeeController(mockEmployeeService)
	response := gin.Default()
	response.PATCH("/employees/:id", sut.UpdateByIdEmployee)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/employees/invalid_id", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should return an error and 422 status if body request contains unprocessable data", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/employees/1", makeUnprocessableCreateAndUpdateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "{\"error\":")
	})

	t.Run("Should return an error and 400 status if body request contains invalid data", func(t *testing.T) {
		testCases := makeInvalidCreateAndUpdateBodiesTestCases()
		for _, tc := range testCases {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, "/employees/1", bytes.NewBuffer([]byte(tc.RequestBody)))
			response.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, tc.ExpectedResponseBody, rr.Body.String())
		}
	})

	t.Run("Should call UpdateById from employee Service with correct values", func(t *testing.T) {
		mockEmployeeService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(makeDBEmployee(), nil).Once()
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(makeDBEmployee(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/employees/1", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		mockEmployeeService.AssertCalled(t, "UpdateById", 1, 568, "Valid_Name", "Valid_Last_Name", 1)
	})

	t.Run("Should return an error and 404 status if UpdateById from employee Service returns an Business Rule error", func(t *testing.T) {
		mockEmployeeService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(employee.Employee{}, &employee.BusinessRuleError{Err: errors.New("any_message")}).Once()
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(employee.Employee{}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/employees/1", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if UpdateById from Employee Service did not returns an custom error", func(t *testing.T) {
		mockEmployeeService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(employee.Employee{}, errors.New("any_message")).Once()
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(employee.Employee{}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/employees/1", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockEmployeeService.On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(makeDBEmployee(), nil).Once()
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(employee.Employee{}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/employees/1", makeValidCreateBody())
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"card_number_id\":568,\"first_name\":\"Valid_Name\",\"Last_name\":\"Valid_Last_Name\",\"warehouse_id\":1}}", rr.Body.String())
	})
}

func TestGetAllEmployees(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockEmployeeService := mocks.NewEmployeeServiceInterface(t)
	sut := employeeController.CreateEmployeeController(mockEmployeeService)
	response := gin.Default()
	response.GET("employees", sut.GetAllEmployee)

	t.Run("Should call GetAll from employees service", func(t *testing.T) {
		mockEmployeeService.On("GetAll").Return([]employee.Employee{makeDBEmployee()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees", nil)
		response.ServeHTTP(rr, req)
	})

	t.Run("Should get 200 status and data on success", func(t *testing.T) {
		mockEmployeeService.On("GetAll").Return([]employee.Employee{makeDBEmployee()}, nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees", nil)
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":[{\"id\":1,\"card_number_id\":568,\"first_name\":\"Valid_Name\",\"Last_name\":\"Valid_Last_Name\",\"warehouse_id\":1}]}", rr.Body.String())
	})
}

func TestGetByIdEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockEmployeeService := mocks.NewEmployeeServiceInterface(t)
	sut := employeeController.CreateEmployeeController(mockEmployeeService)

	response := gin.Default()
	response.GET("/employees/:id", sut.GetByIdEmployee)

	t.Run("Should return an error and 400 status if a invalid id is provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees/invalid_id", nil)
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", rr.Body.String())
	})

	t.Run("Should call GetById from Employee Service with correct id", func(t *testing.T) {
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(makeDBEmployee(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees/1", nil)
		response.ServeHTTP(rr, req)

		mockEmployeeService.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return an error and 404 status if GetById from Employees Service returns not find the correspondent element", func(t *testing.T) {
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(employee.Employee{}, &employee.NoElementInFileError{Err: errors.New("any_message")}).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees/404", nil)
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should return an error and 500 status if GetById from Employees Service returns an error", func(t *testing.T) {
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(employee.Employee{}, errors.New("any_message")).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees/1", nil)
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "{\"error\":\"any_message\"}", rr.Body.String())
	})

	t.Run("Should 200 status and data on success", func(t *testing.T) {
		mockEmployeeService.On("GetById", mock.AnythingOfType("int")).Return(makeDBEmployee(), nil).Once()
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/employees/1", nil)
		response.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "{\"data\":{\"id\":1,\"card_number_id\":568,\"first_name\":\"Valid_Name\",\"Last_name\":\"Valid_Last_Name\",\"warehouse_id\":1}}", rr.Body.String())
	})
}
