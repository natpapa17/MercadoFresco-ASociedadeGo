package adapters_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	makeCreateParams := func() (string, string, string, int, float64) {
		return "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0
	}

	makeWarehouse := func() domain.Warehouse {
		return domain.Warehouse{
			Id:                 1,
			WarehouseCode:      "valid_code",
			Address:            "valid_address",
			Telephone:          "(99) 99999-9999",
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
		}
	}

	makeSut := func() (usecases.WarehouseRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateWarehouseMySQLRepository(db)

		return sut, mock
	}

	t.Run("Should return err if begin transaction fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectBegin().WillReturnError(errors.New("any_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Warehouse{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO warehouse").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute rollback if query fails", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO warehouse").WillReturnError(errors.New("any_error"))
		mock.ExpectRollback()

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Warehouse{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should commit in insert query success", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO warehouse").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if commit fails", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO warehouse").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit().WillReturnError(errors.New("commit_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Warehouse{})
		assert.EqualError(t, err, "commit_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return inserted warehouse on success", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO warehouse").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := sut.Create(makeCreateParams())

		expected := makeWarehouse()

		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	makeGetAllReturn := func() domain.Warehouses {
		return domain.Warehouses{
			domain.Warehouse{
				Id:                 1,
				WarehouseCode:      "valid_code_1",
				Address:            "valid_address_1",
				Telephone:          "valid_phone_1",
				MinimumCapacity:    11,
				MinimumTemperature: 1.8,
			},
			domain.Warehouse{
				Id:                 2,
				WarehouseCode:      "valid_code_2",
				Address:            "valid_address_2",
				Telephone:          "valid_phone_2",
				MinimumCapacity:    12,
				MinimumTemperature: 2.8,
			},
			domain.Warehouse{
				Id:                 3,
				WarehouseCode:      "valid_code_3",
				Address:            "valid_address_3",
				Telephone:          "valid_phone_3",
				MinimumCapacity:    13,
				MinimumTemperature: 3.8,
			},
		}
	}
	makeSut := func() (usecases.WarehouseRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateWarehouseMySQLRepository(db)

		return sut, mock
	}
	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()
		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})
		rows.AddRow(1, "valid_code_1", "valid_address_1", "valid_phone_1", 11, 1.8)
		rows.AddRow(2, "valid_code_2", "valid_address_2", "valid_phone_2", 12, 2.8)
		rows.AddRow(3, "valid_code_3", "valid_address_3", "valid_phone_3", 13, 3.8)
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs().WillReturnRows(rows)

		sut.GetAll()

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs().WillReturnError(errors.New("query_error"))
		result, err := sut.GetAll()

		assert.Equal(t, domain.Warehouses{}, result)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return locality slice on success", func(t *testing.T) {
		sut, mock := makeSut()
		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})
		rows.AddRow(1, "valid_code_1", "valid_address_1", "valid_phone_1", 11, 1.8)
		rows.AddRow(2, "valid_code_2", "valid_address_2", "valid_phone_2", 12, 2.8)
		rows.AddRow(3, "valid_code_3", "valid_address_3", "valid_phone_3", 13, 3.8)
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs().WillReturnRows(rows)

		result, err := sut.GetAll()

		expected := makeGetAllReturn()
		assert.Equal(t, expected, result)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	makeWarehouse := func() domain.Warehouse {
		return domain.Warehouse{
			Id:                 1,
			WarehouseCode:      "valid_code",
			Address:            "valid_address",
			Telephone:          "(99) 99999-9999",
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
		}
	}
	makeSut := func() (usecases.WarehouseRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateWarehouseMySQLRepository(db)

		return sut, mock
	}
	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()

		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})
		rows.AddRow(1, "valid_code_1", "valid_address_1", "valid_phone_1", 11, 1.8)
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs(1).WillReturnRows(rows)

		sut.GetById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return ErrNoElementFound if can't find element in database", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs(1).WillReturnError(sql.ErrNoRows)

		result, err := sut.GetById(1)

		assert.Equal(t, result, domain.Warehouse{})
		assert.Equal(t, err, usecases.ErrNoElementFound)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs(1).WillReturnError(errors.New("query_error"))

		result, err := sut.GetById(1)

		assert.Equal(t, domain.Warehouse{}, result)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an warehouse on success", func(t *testing.T) {
		sut, mock := makeSut()
		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})
		rows.AddRow(1, "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0)
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs(1).WillReturnRows(rows)

		result, err := sut.GetById(1)

		assert.Equal(t, makeWarehouse(), result)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetByWarehouseCode(t *testing.T) {
	makeWarehouse := func() domain.Warehouse {
		return domain.Warehouse{
			Id:                 1,
			WarehouseCode:      "valid_code",
			Address:            "valid_address",
			Telephone:          "(99) 99999-9999",
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
		}
	}
	makeSut := func() (usecases.WarehouseRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateWarehouseMySQLRepository(db)

		return sut, mock
	}
	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()

		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})
		rows.AddRow(1, "valid_code", "valid_address_1", "valid_phone_1", 11, 1.8)
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs("valid_code").WillReturnRows(rows)

		sut.GetByWarehouseCode("valid_code")

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return ErrNoElementFound if can't find element in database", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs("valid_code").WillReturnError(sql.ErrNoRows)

		result, err := sut.GetByWarehouseCode("valid_code")

		assert.Equal(t, result, domain.Warehouse{})
		assert.Equal(t, err, usecases.ErrNoElementFound)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs("valid_code").WillReturnError(errors.New("query_error"))

		result, err := sut.GetByWarehouseCode("valid_code")

		assert.Equal(t, domain.Warehouse{}, result)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an warehouse on success", func(t *testing.T) {
		sut, mock := makeSut()
		rows := sqlmock.NewRows([]string{"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature"})
		rows.AddRow(1, "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0)
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WithArgs("valid_code").WillReturnRows(rows)

		result, err := sut.GetByWarehouseCode("valid_code")

		assert.Equal(t, makeWarehouse(), result)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	makeWarehouse := func() domain.Warehouse {
		return domain.Warehouse{
			Id:                 1,
			WarehouseCode:      "valid_code",
			Address:            "valid_address",
			Telephone:          "(99) 99999-9999",
			MinimumCapacity:    10,
			MinimumTemperature: 5.0,
		}
	}
	makeUpdateParams := func() (int, string, string, string, int, float64) {
		return 1, "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0
	}
	makeSut := func() (usecases.WarehouseRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateWarehouseMySQLRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("UPDATE warehouse SET").WithArgs("valid_code", "valid_address", "(99) 99999-9999", 10, 5.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		sut.UpdateById(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("UPDATE warehouse SET").WithArgs("valid_code", "valid_address", "(99) 99999-9999", 10, 5.0, 1).WillReturnError(errors.New("query_error"))

		result, updateErr := sut.UpdateById(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Equal(t, domain.Warehouse{}, result)
		assert.EqualError(t, updateErr, "query_error")
	})

	t.Run("Should return ErrNoElementFound if element did not exists in db", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("UPDATE warehouse SET").WithArgs("valid_code", "valid_address", "(99) 99999-9999", 10, 5.0, 1).WillReturnResult(sqlmock.NewResult(1, 0))
		mock.ExpectQuery("SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse").WillReturnError(sql.ErrNoRows)

		result, updateErr := sut.UpdateById(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Equal(t, domain.Warehouse{}, result)
		assert.Equal(t, updateErr, usecases.ErrNoElementFound)
	})

	t.Run("Should return updated warehouse on success", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("UPDATE warehouse SET").WithArgs("valid_code", "valid_address", "(99) 99999-9999", 10, 5.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		result, updatedErr := sut.UpdateById(makeUpdateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Equal(t, makeWarehouse(), result)
		assert.Nil(t, updatedErr)
	})
}

func TestDeleteById(t *testing.T) {
	makeSut := func() (usecases.WarehouseRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateWarehouseMySQLRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("DELETE FROM warehouse").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		sut.DeleteById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("DELETE FROM warehouse").WithArgs(1).WillReturnError(errors.New("query_error"))

		deleteErr := sut.DeleteById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.EqualError(t, deleteErr, "query_error")
	})

	t.Run("Should return ErrNoElementFound if element did not exists in db", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("DELETE FROM warehouse").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 0))

		deleteErr := sut.DeleteById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Equal(t, deleteErr, usecases.ErrNoElementFound)
	})

	t.Run("Should return nil on success", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectExec("DELETE FROM warehouse").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		deleteErr := sut.DeleteById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)

		assert.Nil(t, deleteErr)
	})
}
