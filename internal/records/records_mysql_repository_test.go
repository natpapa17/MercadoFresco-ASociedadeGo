package records_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records"
	"github.com/stretchr/testify/assert"
)

func TestGetRecordsPerProduct(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sut := records.NewMysqlRepository(db)

	t.Run("Should execute correct query in database", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"count"})
		rows.AddRow(1)
		mock.ExpectQuery("SELECT COUNT(.+) FROM records").WithArgs(1).WillReturnRows(rows)

		sut.GetRecordsPerProduct(1)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		mock.ExpectQuery("SELECT COUNT(.+) FROM records").WithArgs(1).WillReturnError(errors.New("query_error"))

		result, err := sut.GetRecordsPerProduct(1)

		assert.Equal(t, result, 0)
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an integer on success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"count"})
		rows.AddRow(1)
		mock.ExpectQuery("SELECT COUNT(.+) FROM records").WithArgs(1).WillReturnRows(rows)

		result, err := sut.GetRecordsPerProduct(1)

		assert.Equal(t, result, 1)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetByProductId(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sut := records.NewMysqlRepository(db)

	t.Run("Shoul execute correct query in database", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"})
		rows.AddRow(1, "2000-10-10", 5, 10, 1)
		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM records").WithArgs(1).WillReturnRows(rows)

		sut.GetByProductId(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return product not found if can't find element in database", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM records").WithArgs(1).WillReturnError(errors.New("product not found"))

		result, err := sut.GetByProductId(1)

		assert.Equal(t, result, records.Records{})
		assert.EqualError(t, err, "product not found")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an error if query fails", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM records").WithArgs(1).WillReturnError(errors.New("query_error"))

		result, err := sut.GetByProductId(1)

		assert.Equal(t, result, records.Records(records.Records{Id: 0, Last_Update_Date: "", Purchase_Price: 0, Sale_Price: 0, Product_Id: 0}))
		assert.EqualError(t, err, "query_error")

		sut.GetByProductId(1)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an record on success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"})
		rows.AddRow(1, "2000-10-10", 5, 10, 1)
		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM records").WithArgs(1).WillReturnRows(rows)

		result, err := sut.GetByProductId(1)

		expected := records.Records{
			Id:               1,
			Last_Update_Date: "2000-10-10",
			Purchase_Price:   5,
			Sale_Price:       10,
			Product_Id:       1,
		}
		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestCreate(t *testing.T) {
	makeCreateParams := func() (string, int, int, int) {
		return "2000-10-10", 5, 10, 1
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sut := records.NewMysqlRepository(db)

	t.Run("Should execute correct query in database", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO records").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should execute rollback if query fails", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO records").WillReturnError(errors.New("any_error"))
		mock.ExpectRollback()

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, result, records.Records{})
		assert.EqualError(t, err, "any_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should commit insert query success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO records").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		sut.Create(makeCreateParams())

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return Create Ok", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO records").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := sut.Create(makeCreateParams())

		expected := records.Records{
			Id:               1,
			Last_Update_Date: "2000-10-10",
			Purchase_Price:   5,
			Sale_Price:       10,
			Product_Id:       1,
		}
		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}
