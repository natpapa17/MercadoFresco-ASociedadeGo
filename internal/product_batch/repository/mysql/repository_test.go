package mysql

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := NewMySQLRepository(db)

	t.Run("create_ok", func(t *testing.T) {
		mock.
			ExpectExec("INSERT INTO product_batch").
			WithArgs(1, 111, 200, 20, "2022-04-04", 20, "2022-04-04", 10, 5, 1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		pb, err := repo.Add(context.Background(), 1, 111, 200, 20, "2022-04-04", 20, "2022-04-04", 10, 5, 1, 1)

		epb := product_batch.ProductBatch{
			ID:                 1,
			BatchNumber:        111,
			CurrentQuantity:    200,
			CurrentTemperature: 20,
			DueDate:            "2022-04-04",
			InitialQuantity:    20,
			ManufacturingDate:  "2022-04-04",
			ManufacturingHour:  10,
			MinimumTemperature: 5,
			ProductID:          1,
			SectionID:          1,
		}

		assert.Nil(t, err)
		assert.Equal(t, epb, pb)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)

	})

	t.Run("create_fail", func(t *testing.T) {
		mock.
			ExpectExec("INSERT INTO product_batch").
			WithArgs(1, 111, 200, 20, "2022-04-04", 20, "2022-04-04", 10, 5, 1, 1).
			WillReturnError(fmt.Errorf("error"))

		pb, err := repo.Add(context.Background(), 1, 111, 200, 20, "2022-04-04", 20, "2022-04-04", 10, 5, 1, 1)

		assert.Error(t, err)
		assert.Equal(t, product_batch.ProductBatch{}, pb)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

}
