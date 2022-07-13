package adapters_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	makeCreateParams := func() (string, string, string, string, int) {
		return "valid_cid", "valid_name", "valid_address", "valid_phone", 1
	}

	makeSut := func() (usecases.CarrierRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateCarrierMySQLRepository(db)

		return sut, mock
	}

	t.Run("Should return err if begin transaction fails", func(t *testing.T) {
		sut, mock := makeSut()
		mock.ExpectBegin().WillReturnError(errors.New("any_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Carrier{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO carrier").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute rollback if query fails", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO carrier").WillReturnError(errors.New("any_error"))
		mock.ExpectRollback()

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Carrier{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should commit in insert query success", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO carrier").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if commit fails", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO carrier").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit().WillReturnError(errors.New("commit_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, domain.Carrier{})
		assert.EqualError(t, err, "commit_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return inserted carrier on success", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO carrier").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := sut.Create(makeCreateParams())

		expected := domain.Carrier{
			Id:          1,
			Cid:         "valid_cid",
			CompanyName: "valid_name",
			Address:     "valid_address",
			Telephone:   "valid_phone",
			LocalityId:  1,
		}

		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetNumberOfCarriersPerLocality(t *testing.T) {
	makeSut := func() (usecases.CarrierRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateCarrierMySQLRepository(db)

		return sut, mock
	}

	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()
		rows := sqlmock.NewRows([]string{"count"})
		rows.AddRow(1)
		mock.ExpectQuery("SELECT COUNT(.+) FROM carrier").WithArgs(1).WillReturnRows(rows)

		sut.GetNumberOfCarriersPerLocality(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectQuery("SELECT COUNT(.+) FROM carrier").WithArgs(1).WillReturnError(errors.New("query_error"))

		result, err := sut.GetNumberOfCarriersPerLocality(1)

		assert.Equal(t, result, 0)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an integer on success", func(t *testing.T) {
		sut, mock := makeSut()

		rows := sqlmock.NewRows([]string{"count"})
		rows.AddRow(1)
		mock.ExpectQuery("SELECT COUNT(.+) FROM carrier").WithArgs(1).WillReturnRows(rows)

		result, err := sut.GetNumberOfCarriersPerLocality(1)

		assert.Equal(t, result, 1)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetByCid(t *testing.T) {
	makeSut := func() (usecases.CarrierRepository, sqlmock.Sqlmock) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sut := adapters.CreateCarrierMySQLRepository(db)

		return sut, mock
	}
	t.Run("Should execute correct query in database", func(t *testing.T) {
		sut, mock := makeSut()

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"})
		rows.AddRow(1, "valid_cid", "valid_name", "valid_address", "valid_phone", 1)
		mock.ExpectQuery("SELECT id, cid, company_name, address, telephone, locality_id FROM carrier").WithArgs("valid_cid").WillReturnRows(rows)

		sut.GetByCid("valid_cid")

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return ErrNoElementFound if can't find element in database", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectQuery("SELECT id, cid, company_name, address, telephone, locality_id FROM carrier").WithArgs("valid_cid").WillReturnError(sql.ErrNoRows)

		result, err := sut.GetByCid("valid_cid")

		assert.Equal(t, result, domain.Carrier{})
		assert.Equal(t, err, usecases.ErrNoElementFound)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		sut, mock := makeSut()

		mock.ExpectQuery("SELECT id, cid, company_name, address, telephone, locality_id FROM carrier").WithArgs("valid_cid").WillReturnError(errors.New("query_error"))

		result, err := sut.GetByCid("valid_cid")

		assert.Equal(t, result, domain.Carrier{})
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an carrier on success", func(t *testing.T) {
		sut, mock := makeSut()

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"})
		rows.AddRow(1, "valid_cid", "valid_name", "valid_address", "valid_phone", 1)
		mock.ExpectQuery("SELECT id, cid, company_name, address, telephone, locality_id FROM carrier").WithArgs("valid_cid").WillReturnRows(rows)

		result, err := sut.GetByCid("valid_cid")

		expected := domain.Carrier{
			Id:          1,
			Cid:         "valid_cid",
			CompanyName: "valid_name",
			Address:     "valid_address",
			Telephone:   "valid_phone",
			LocalityId:  1,
		}
		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}
