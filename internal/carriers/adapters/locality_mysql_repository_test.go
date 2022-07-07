package adapters_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sut := adapters.CreateLocalityMySQLRepository(db)
	t.Run("Should execute correct query in database", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "province_id"})
		rows.AddRow(1, "valid_name", 1)
		mock.ExpectQuery("SELECT id, name, province_id FROM locality").WithArgs(1).WillReturnRows(rows)

		sut.GetById(1)

		err := mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return ErrNoElementFound if can't find element in database", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, province_id FROM locality").WithArgs(1).WillReturnError(sql.ErrNoRows)

		result, err := sut.GetById(1)

		assert.Equal(t, result, domain.Locality{})
		assert.Equal(t, err, usecases.ErrNoElementFound)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return error if query fails", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, province_id FROM locality").WithArgs(1).WillReturnError(errors.New("query_error"))

		result, err := sut.GetById(1)

		assert.Equal(t, result, domain.Locality{})
		assert.EqualError(t, err, "query_error")

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("Should return an locality on success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "province_id"})
		rows.AddRow(1, "valid_name", 1)
		mock.ExpectQuery("SELECT id, name, province_id FROM locality").WithArgs(1).WillReturnRows(rows)

		result, err := sut.GetById(1)

		expected := domain.Locality{
			Id:         1,
			Name:       "valid_name",
			ProvinceId: 1,
		}
		assert.Equal(t, result, expected)
		assert.Nil(t, err)

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}