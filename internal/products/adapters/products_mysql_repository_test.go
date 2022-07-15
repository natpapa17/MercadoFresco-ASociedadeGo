package adapters_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	makeGetAllReturn := func() domain.Products {
		return domain.Products{
			domain.Product{
				Id:                               1,
				Product_Code:                     "PROD01",
				Description:                      "valid_description_1",
				Width:                            1.0,
				Height:                           1.0,
				Length:                           1.0,
				Net_Weight:                       1.0,
				Expiration_Rate:                  1,
				Recommended_Freezing_Temperature: 1.0,
				Freezing_Rate:                    1,
				Product_Type_Id:                  1,
				Seller_Id:                        1,
			},
			domain.Product{
				Id:                               2,
				Product_Code:                     "PROD02",
				Description:                      "valid_description_2",
				Width:                            2.0,
				Height:                           2.0,
				Length:                           2.0,
				Net_Weight:                       2.0,
				Expiration_Rate:                  2,
				Recommended_Freezing_Temperature: 2.0,
				Freezing_Rate:                    2,
				Product_Type_Id:                  2,
				Seller_Id:                        2,
			},
			domain.Product{
				Id:                               3,
				Product_Code:                     "PROD03",
				Description:                      "valid_description_3",
				Width:                            3.0,
				Height:                           3.0,
				Length:                           3.0,
				Net_Weight:                       3.0,
				Expiration_Rate:                  3,
				Recommended_Freezing_Temperature: 3.0,
				Freezing_Rate:                    3,
				Product_Type_Id:                  3,
				Seller_Id:                        3,
			},
		}
	}

	makeRepository := func() (usecases.RepositoryProduct, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.NewProductMysqlRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeRepository()
		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		rows.AddRow(1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
		rows.AddRow(2, "PROD02", "valid_description_2", 2.0, 2.0, 2.0, 2.0, 2, 2.0, 2, 2, 2)
		rows.AddRow(3, "PROD03", "valid_description_3", 3.0, 3.0, 3.0, 3.0, 3, 3.0, 3, 3, 3)
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs().WillReturnRows(rows)

		sut.GetAll()

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs().WillReturnError(errors.New("query_error"))
		result, err := sut.GetAll()

		assert.Equal(t, domain.Products{}, result)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return locality slice on success", func(t *testing.T) {
		sut, mock := makeRepository()
		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		rows.AddRow(1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
		rows.AddRow(2, "PROD02", "valid_description_2", 2.0, 2.0, 2.0, 2.0, 2, 2.0, 2, 2, 2)
		rows.AddRow(3, "PROD03", "valid_description_3", 3.0, 3.0, 3.0, 3.0, 3, 3.0, 3, 3, 3)
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs().WillReturnRows(rows)

		result, err := sut.GetAll()

		expected := makeGetAllReturn()
		assert.Equal(t, expected, result)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	makeProduct := func() domain.Product {
		return domain.Product{

			Id:                               1,
			Product_Code:                     "PROD01",
			Description:                      "valid_description_1",
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

	makeRepository := func() (usecases.RepositoryProduct, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.NewProductMysqlRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeRepository()

		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		rows.AddRow(1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs(1).WillReturnRows(rows)

		sut.GetById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs(1).WillReturnError(errors.New("query_error"))

		result, err := sut.GetById(1)

		assert.Equal(t, domain.Product{}, result)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an product on success", func(t *testing.T) {
		sut, mock := makeRepository()
		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		rows.AddRow(1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs(1).WillReturnRows(rows)

		result, err := sut.GetById(1)

		assert.Equal(t, makeProduct(), result)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetByCode(t *testing.T) {
	makeProduct := func() domain.Product {
		return domain.Product{

			Id:                               1,
			Product_Code:                     "PROD01",
			Description:                      "valid_description_1",
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

	makeRepository := func() (usecases.RepositoryProduct, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.NewProductMysqlRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeRepository()

		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		rows.AddRow(1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs("PROD01").WillReturnRows(rows)

		sut.GetByCode("PROD01")

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs("PROD01").WillReturnError(errors.New("query_error"))

		result, err := sut.GetByCode("PROD01")

		assert.Equal(t, domain.Product{}, result)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an Product on success", func(t *testing.T) {
		sut, mock := makeRepository()
		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		rows.AddRow(1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1)
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product").WithArgs("PROD01").WillReturnRows(rows)

		result, err := sut.GetByCode("PROD01")

		assert.Equal(t, makeProduct(), result)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestCreate(t *testing.T) {
	makeCreateParams := func() (string, string, float64, float64, float64, float64, int, float64, int, int, int) {
		return "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1
	}

	makeProduct := func() domain.Product {
		return domain.Product{

			Id:                               1,
			Product_Code:                     "PROD01",
			Description:                      "valid_description_1",
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

	makeRepository := func() (usecases.RepositoryProduct, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.NewProductMysqlRepository(db)

		return sut, mock
	}

	t.Run("Should return err if begin transaction fails", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectBegin().WillReturnError(errors.New("any_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Product{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO product").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute rollback if query fails", func(t *testing.T) {
		sut, mock := makeRepository()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO product").WillReturnError(errors.New("any_error"))
		mock.ExpectRollback()

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Product{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should commit in insert query success", func(t *testing.T) {
		sut, mock := makeRepository()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO product").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if commit fails", func(t *testing.T) {
		sut, mock := makeRepository()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO product").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit().WillReturnError(errors.New("commit_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Product{})
		assert.EqualError(t, err, "commit_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return inserted warehouse on success", func(t *testing.T) {
		sut, mock := makeRepository()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO product").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := sut.Create(makeCreateParams())

		expected := makeProduct()

		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	makeProduct := func() domain.Product {
		return domain.Product{
			Id:                               1,
			Product_Code:                     "PROD01",
			Description:                      "valid_description_1",
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

	makeUpdateParams := func() (int, string, string, float64, float64, float64, float64, int, float64, int, int, int) {
		return 1, "PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1
	}

	makeRepository := func() (usecases.RepositoryProduct, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.NewProductMysqlRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectExec("UPDATE product SET").WithArgs("PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		sut.Update(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectExec("UPDATE product SET").WithArgs("PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1, 1).WillReturnError(errors.New("query_error"))

		result, updateErr := sut.Update(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Equal(t, domain.Product{}, result)
		assert.EqualError(t, updateErr, "query_error")
	})

	t.Run("Should return updated warehouse on success", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectExec("UPDATE product SET").WithArgs("PROD01", "valid_description_1", 1.0, 1.0, 1.0, 1.0, 1, 1.0, 1, 1, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		result, updatedErr := sut.Update(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Equal(t, makeProduct(), result)
		assert.Nil(t, updatedErr)
	})
}

func TestDeleteById(t *testing.T) {
	makeRepository := func() (usecases.RepositoryProduct, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.NewProductMysqlRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectExec("DELETE FROM product").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		sut.Delete(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectExec("DELETE FROM product").WithArgs(1).WillReturnError(errors.New("query_error"))

		deleteErr := sut.Delete(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.EqualError(t, deleteErr, "query_error")
	})

	t.Run("Should return nil on success", func(t *testing.T) {
		sut, mock := makeRepository()
		mock.ExpectExec("DELETE FROM product").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		deleteErr := sut.Delete(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Nil(t, deleteErr)
	})
}
