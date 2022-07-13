package product_batch_test

import (
	"context"
	"errors"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func createProductReport(sectionID int, sectionNumber int, productsCount int) product_batch.ProductsReport {
	return product_batch.ProductsReport{
		SectionID:     sectionID,
		SectionNumber: sectionNumber,
		ProductsCount: productsCount,
	}
}

func createProductBatch(batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature int, productID int, sectionID int) product_batch.ProductBatch {
	return product_batch.ProductBatch{
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: minimumTemperature,
		ProductID:          productID,
		SectionID:          sectionID,
	}
}

func createProductBatchWithId(id int, batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature int, productID int, sectionID int) product_batch.ProductBatch {
	return product_batch.ProductBatch{
		ID:                 id,
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: minimumTemperature,
		ProductID:          productID,
		SectionID:          sectionID,
	}
}

func TestGetById(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := product_batch.NewService(repo)
	ctx := context.Background()

	t.Run("get_by_id_ok", func(t *testing.T) {
		repo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Return([]product_batch.ProductsReport{createProductReport(1, 23, 200)}, nil).
			Once()

		report, err := serv.GetById(ctx, 1)

		assert.NoError(t, err)
		assert.Equal(t, []product_batch.ProductsReport{{1, 23, 200}}, report)
	})

	t.Run("get_by_id_fail", func(t *testing.T) {
		repo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Return([]product_batch.ProductsReport{}, errors.New("not found")).
			Once()
	})

	_, err := serv.GetById(ctx, 1)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := product_batch.NewService(repo)
	ctx := context.Background()

	t.Run("add_ok", func(t *testing.T) {
		repo.
			On("LastID", mock.Anything).
			Return(0, nil).
			Once()

		repo.
			On("HasBatchNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(false, nil).
			Once()

		repo.
			On("Add", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(createProductBatchWithId(1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1), nil).
			Once()

		pb, err := serv.Add(ctx, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1)
		assert.NoError(t, err)
		assert.Equal(t, product_batch.ProductBatch{1, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1}, pb)
	})

	t.Run("add_fail", func(t *testing.T) {
		repo.
			On("HasBatchNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(true, nil).
			Once()

		_, err := serv.Add(ctx, 111, 200, 20, "2022-04-04", 10, "2020-04-04", 10, 5, 1, 1)
		assert.Error(t, err)
	})

}

func TestLastID(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := product_batch.NewService(repo)
	ctx := context.Background()

	t.Run("last_id_ok", func(t *testing.T) {
		repo.
			On("LastID", mock.Anything).
			Return(1, nil).
			Once()

		id, err := serv.LastID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("last_id_fail", func(t *testing.T) {
		repo.
			On("LastID", mock.Anything).
			Return(0, errors.New("error")).
			Once()

		_, err := serv.LastID(ctx)
		assert.Error(t, err)
	})
}

func TestHasBatchNumber(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := product_batch.NewService(repo)
	ctx := context.Background()

	t.Run("has_batch_number_ok", func(t *testing.T) {
		repo.
			On("HasBatchNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(false, nil).
			Once()

		has, err := serv.HasBatchNumber(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, false, has)
	})

	t.Run("has_batch_number_fail", func(t *testing.T) {
		repo.
			On("HasBatchNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(true, errors.New("error")).
			Once()

		_, err := serv.HasBatchNumber(ctx, 1)
		assert.Error(t, err)
	})
}
