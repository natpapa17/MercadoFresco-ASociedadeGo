package section_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	controller "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/section"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createValidBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
		"section_number": 99,
		"current_temperature": 1,
		"minimum_temperature": 1,
		"current_capacity": 2,
		"minimum_capacity": 1,
		"maximum_capacity": 4,
		"warehouse_id": 1,
		"product_type_id": 1
	}`))
}

func createInvalidBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(`
	{
		"current_temperature": 1,
		"minimum_temperature": 1,
		"current_capacity": 2,
		"minimum_capacity": 1,
		"maximum_capacity": 4,
		"warehouse_id": 1,
		"product_type_id": 1
	}`))
}

func createValidJSONWithParams(id int, section int, currentTemperature float32, minimumTemperature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) string {
	return fmt.Sprintf("{\"id\":%d,\"section_number\":%d,\"current_temperature\":%0.f,\"minimum_temperature\":%0.f,\"current_capacity\":%d,\"minimum_capacity\":%d,\"maximum_capacity\":%d,\"warehouse_id\":%d,\"product_type_id\":%d}", id, section, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID)
}

func createValidSectionWithParams(id int, section int, currentTemperature float32, minimumTemperature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) sections.Section {
	return sections.Section{
		ID:                 id,
		SectionNumber:      section,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}
}

func TestGetAll(t *testing.T) {

	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	sc := controller.NewSection(mockService)

	r := gin.Default()
	r.GET("/section", sc.GetAll())

	t.Run("find_all", func(t *testing.T) {
		mockService.
			On("GetAll", mock.Anything).
			Return([]sections.Section{createValidSectionWithParams(1, 1, 1, 1, 1, 1, 1, 1, 1), createValidSectionWithParams(2, 2, 1, 1, 1, 1, 1, 1, 1), createValidSectionWithParams(3, 3, 1, 1, 1, 1, 1, 1, 1)}, nil).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/section", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestAdd(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	sc := controller.NewSection(mockService)

	r := gin.Default()
	r.POST("/section", sc.Add())

	t.Run("create_ok", func(t *testing.T) {
		mockService.
			On("Add", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(createValidSectionWithParams(1, 99, 1, 1, 2, 1, 4, 1, 1), nil).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/section", createValidBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, createValidJSONWithParams(1, 99, 1, 1, 2, 1, 4, 1, 1), rr.Body.String())

	})

	t.Run("create_fail", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/section", createInvalidBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.JSONEq(t, `{"error":"Key: 'SectionRequest.SectionNumber' Error:Field validation for 'SectionNumber' failed on the 'required' tag"}`, rr.Body.String())

	})

	t.Run("create_conflict", func(t *testing.T) {
		mockService.
			On("Add", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(sections.Section{}, errors.New("section already exists")).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/section", createValidBody())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})
}

func TestGetById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	sc := controller.NewSection(mockService)

	r := gin.Default()
	r.POST("/section/:id", sc.GetById())

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockService.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Return(sections.Section{}, errors.New("Id not found.")).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/section/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.JSONEq(t, `{"error":"Id not found."}`, rr.Body.String())
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		mockService.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Return(createValidSectionWithParams(1, 1, 1, 1, 2, 1, 4, 1, 1), nil).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/section/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, createValidJSONWithParams(1, 1, 1, 1, 2, 1, 4, 1, 1), rr.Body.String())
	})
}

func TestUpdateById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	sc := controller.NewSection(mockService)

	r := gin.Default()
	r.PATCH("/section/:id", sc.UpdateById())

	t.Run("update_ok", func(t *testing.T) {
		mockService.
			On("UpdateById", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("sections.Section")).
			Return(createValidSectionWithParams(1, 1, 1, 1, 5, 1, 10, 1, 1), nil).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/section/1", bytes.NewBuffer([]byte(`{"current_capacity": 5, "maximum_capacity": 10}`)))
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, createValidJSONWithParams(1, 1, 1, 1, 5, 1, 10, 1, 1), rr.Body.String())
	})

	t.Run("update_non_existent", func(t *testing.T) {
		mockService.
			On("UpdateById", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("sections.Section")).
			Return(sections.Section{}, errors.New("inexistent section")).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/section/1", bytes.NewBuffer([]byte(`{"current_capacity": 5, "maximum_capacity": 10}`)))
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.JSONEq(t, `{"error":"inexistent section"}`, rr.Body.String())
	})
}

func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	sc := controller.NewSection(mockService)

	r := gin.Default()
	r.DELETE("/section/:id", sc.Delete())

	t.Run("delete_non_existent", func(t *testing.T) {
		mockService.
			On("Delete", mock.Anything, mock.AnythingOfType("int")).
			Return(errors.New("Id not found")).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/section/9", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "{\"error\":\"Id not found\"}", rr.Body.String())
	})

	t.Run("update_ok", func(t *testing.T) {
		mockService.
			On("Delete", mock.Anything, mock.AnythingOfType("int")).
			Return(nil).
			Once()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/section/1", nil)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		assert.Equal(t, "", rr.Body.String())
	})
}
