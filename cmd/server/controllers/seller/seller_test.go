package controllers_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/seller"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func UpdateBody() *bytes.Buffer{
	return bytes.NewBuffer(([]byte(`
	{
		"Id": 1,
		"Cid": 1,
		"CompanyName": "None",
		"Address": "rua sem nome",
		"Telephone": "000000"
	}
	`)))
}

func validSeller() *bytes.Buffer{
	return bytes.NewBuffer(([]byte(`
	{
		"Id": 1,
		"Cid": 1,
		"CompanyName": "None",
		"Address": "rua sem nome",
		"Telephone": "000000"
	}
	`)))
}

func InvalidSeller() *bytes.Buffer{
	return bytes.NewBuffer(([]byte(`
	{
		"Id": 1,
		"Cid":  ,
		"CompanyName": "None",
		"Address": "rua sem nome",
		"Telephone": "000000"
	}
	`)))
}

func dbSeller()sellers.Seller{
	return sellers.Seller{
		Id:    1,
			Cid:  1,
			CompanyName:  "None",
			Address: "none",
			Telephone: "00000",
	}
}

func ValidSellerWithParams(Id, Cid int, CompanyName, Address, Telephone string) sellers.Seller{
	return sellers.Seller{
		Id : Id,
		Cid : Cid,
		CompanyName: CompanyName,
		Address: Address,
		Telephone: Telephone,
	}
}



func TestGetById( t *testing.T){
	gin.SetMode(gin.TestMode)
	service := mocks.NewService(t)
	sellerController := controllers.NewSeller(service)

	r:=gin.Default()
	r.GET("/sellers/:id", sellerController.GetByIdSeller() )

	t.Run("Return the seller whit the specified id", func(t *testing.T){
		service.On("GetById", mock.AnythingOfType("int")).Return(dbSeller(), nil).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/sellers/1", nil)

		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		//assert.Equal(t,"{\"data\": [{\"Id\": 1, \"Cid\": 1, \"CompanyName\": \"None\", \"Address\": \"none\", \"Telephone\": \"00000\"}]}",response.Body.String() )

	})

	
	t.Run(" return an error if the id doesn't exists", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/sellers/invalid_id", nil)
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"invalid id\"}", response.Body.String())
	})

	t.Run("return an erro if the id is not found", func (t *testing.T){
		service.On("GetById", int(999)).Return(sellers.Seller{}, fmt.Errorf("seller not found")).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/sellers/999", nil)
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "{\"error\":\"can't find element with this id\"}", response.Body.String())
	})

}

func TestGetAllController(t *testing.T){
	gin.SetMode(gin.TestMode)
	service := mocks.NewService(t)
	sellerController := controllers.NewSeller(service)

	r:=gin.Default()
	r.GET("/sellers", sellerController.GetAll() )

	t.Run("return all sellers", func(t *testing.T){
		s := sellers.Seller{
			Id:    1,
			Cid:  1,
			CompanyName:  "None",
			Address: "none",
			Telephone: "00000",
		}

		sList := make([]sellers.Seller, 0)
		sList = append(sList, s)
		
		service.On("GetAll").Return(sList, nil ).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/sellers", nil)

		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		//assert.Equal(t,"{\"data\": [{\"Id\": 1, \"Cid\": 1, \"CompanyName\": \"None\", \"Address\": \"none\", \"Telephone\": \"00000\"}]}",response.Body.String() )
		
	})
	t.Run("return an error if GetAll returns an error", func(t *testing.T) {
		service.On("GetAll").Return([]sellers.Seller{}, errors.New("any_message")).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/sellers", nil)
		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "{\"error\":\"internal server error\"}", response.Body.String())
	})





}



func TestDelete(t*testing.T){
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	service := controllers.NewSeller(mockService)
	r := gin.Default()

	r.DELETE("/sellers/:id", service.Delete())

	t.Run("delete the seller with the specified id", func(t *testing.T){
		mockService.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodDelete, "/sellers/1", nil)
		r.ServeHTTP(response, request)
		mockService.AssertCalled(t, "Delete", 1)
	})

	t.Run("return an error if the specified id not exists", func(t *testing.T){
		
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodDelete, "/sellers/id", nil)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("return an error if the specified id not exists", func(t *testing.T){
		mockService.On("Delete", int(999)).Return( fmt.Errorf("seller not found")).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodDelete, "/sellers/999", nil)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestStoreController( t *testing.T){
	
	gin.SetMode(gin.TestMode)
	mockService := mocks.NewService(t)
	sellerController := controllers.NewSeller(mockService)

	r := gin.Default()
	r.POST("/sellers", sellerController.Store())
	expectSeller := sellers.Seller{
		Id:    1,
		Cid:  1,
		CompanyName:  "None",
		Address: "none",
		Telephone: "00000",
	}
	t.Run("Create a new seller", func(t*testing.T){
		mockService.On("Store", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(expectSeller, nil).Once()
		request := httptest.NewRecorder()
		response, _ := http.NewRequest(http.MethodPost, "/sellers",validSeller())
		r.ServeHTTP(request, response)
		assert.Equal(t, http.StatusOK, request.Code)
	})
	t.Run("if any any data is missing, return an error", func(t *testing.T){
		
		request := httptest.NewRecorder()
		response, _ := http.NewRequest(http.MethodPost, "/sellers", InvalidSeller())
		r.ServeHTTP(request, response)
		assert.Equal(t, http.StatusBadRequest, request.Code)
	})



	
}


func TestUpdate(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	service := controllers.NewSeller(mockService)
	r := gin.Default()
	expectSeller := sellers.Seller{
		Id:    1,
		Cid:  1,
		CompanyName:  "None",
		Address: "none",
		Telephone: "00000",
	}
	r.PATCH("/sellers/:id", service.Update())
	t.Run("return the seller with the updated data", func(t *testing.T){
		mockService.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(expectSeller, nil).Once()
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPatch, "/sellers/1", UpdateBody())
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)

	})
	
	t.Run("if any any data is missing, return an error", func(t *testing.T){
		
		request := httptest.NewRecorder()
		response, _ := http.NewRequest(http.MethodPost, "/sellers/1", InvalidSeller())
		r.ServeHTTP(request, response)
		assert.Equal(t, http.StatusNotFound, request.Code)
	})


}
