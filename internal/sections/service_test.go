package sections_test

import (
	"context"
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createSectionParams() (int, float32, float32, int, int, int, int, int) {
	return 1, 1.0, 1.0, 1, 1, 1, 1, 1
}

func createSectionParamsWithContext() (context.Context, int, float32, float32, int, int, int, int, int) {
	return context.Background(), 1, 1.0, 1.0, 1, 1, 1, 1, 1
}

func createSectionParamsUpdated() (int, float32, float32, int, int, int, int, int) {
	return 1, 3.0, 1.0, 3, 1, 4, 1, 1
}

func createSectionFromParams(sectionNumber int, currentTemperature float32, minimumTemperature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) sections.Section {
	return sections.Section{
		ID:                 1,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}
}

func createSectionFromParamsWithID(id int, sectionNumber int, currentTemperature float32, minimumTemperature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) sections.Section {
	return sections.Section{
		ID:                 id,
		SectionNumber:      sectionNumber,
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
	repo := mocks.NewRepository(t)
	serv := sections.NewService(repo)
	ctx := context.Background()

	t.Run("find_all", func(t *testing.T) {
		repo.
			On("GetAll", mock.Anything).
			Return([]sections.Section{
				sections.Section{
					ID:                 1,
					SectionNumber:      1,
					CurrentTemperature: 1.0,
					MinimumTemperature: 1.0,
					CurrentCapacity:    1,
					MinimumCapacity:    1,
					MaximumCapacity:    1,
					WarehouseID:        1,
					ProductTypeID:      1,
				},
				sections.Section{
					ID:                 2,
					SectionNumber:      2,
					CurrentTemperature: 1.0,
					MinimumTemperature: 1.0,
					CurrentCapacity:    1,
					MinimumCapacity:    1,
					MaximumCapacity:    1,
					WarehouseID:        1,
					ProductTypeID:      1,
				},
			}, nil)

		sects, _ := serv.GetAll(ctx)

		expected := []sections.Section{createSectionFromParams(createSectionParams()), createSectionFromParamsWithID(2, 2, 1.0, 1.0, 1, 1, 1, 1, 1)}

		assert.Equal(t, sects, expected)
	})
}

func TestGetById(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := sections.NewService(repo)
	ctx := context.Background()

	t.Run("find_by_id_existent", func(t *testing.T) {
		repo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Return(createSectionFromParams(createSectionParams()), nil).
			Once()

		sect, _ := serv.GetById(ctx, 1)

		assert.Equal(t, createSectionFromParams(1, 1.0, 1.0, 1, 1, 1, 1, 1), sect)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		repo.
			On("GetById", mock.Anything, mock.AnythingOfType("int")).
			Return(sections.Section{}, errors.New("Id not found."))

		_, err := serv.GetById(ctx, 1)

		assert.EqualError(t, err, "Id not found.")
	})
}

func TestAdd(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := sections.NewService(repo)

	t.Run("create_ok", func(t *testing.T) {
		repo.
			On("HasSectionNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(false, nil)

		repo.
			On("LastID", mock.Anything).
			Return(0, nil)

		repo.
			On("Add", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(createSectionFromParams(createSectionParams()), nil).
			Once()

		sect, _ := serv.Add(createSectionParamsWithContext())

		assert.Equal(t, createSectionFromParams(createSectionParams()), sect)
	})

	t.Run("create_conflict", func(t *testing.T) {
		repo.
			On("HasSectionNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(true, errors.New("section already exists"))

		repo.
			On("Add", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(sections.Section{}, errors.New("section already exists")).
			Once()

		_, err := serv.Add(createSectionParamsWithContext())

		assert.EqualError(t, err, "section already exists")

	})
}

func TestUpdateById(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := sections.NewService(repo)
	ctx := context.Background()

	t.Run("update_existent", func(t *testing.T) {
		repo.
			On("HasSectionNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(true, nil).
			Once()

		repo.
			On("UpdateById", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("sections.Section")).
			Return(createSectionFromParams(createSectionParamsUpdated()), nil).
			Once()

		sect, _ := serv.UpdateById(ctx, 1, sections.Section{1, 1, 3.0, 1.0, 3, 1, 4, 1, 1})

		assert.Equal(t, createSectionFromParams(createSectionParamsUpdated()), sect)
	})

	t.Run("update_inexistent", func(t *testing.T) {
		repo.
			On("HasSectionNumber", mock.Anything, mock.AnythingOfType("int")).
			Return(false, nil)

		repo.
			On("UpdateById", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("sections.Section")).
			Return(sections.Section{}, errors.New("inexistent section"))

		_, err := serv.UpdateById(ctx, 3, sections.Section{1, 1, 3.0, 1.0, 3, 1, 4, 1, 1})

		assert.EqualError(t, err, "inexistent section")
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewRepository(t)
	serv := sections.NewService(repo)
	ctx := context.Background()

	t.Run("delete_non_existent", func(t *testing.T) {
		repo.
			On("Delete", mock.Anything, mock.AnythingOfType("int")).
			Return(errors.New("Id not found")).
			Once()

		err := serv.Delete(ctx, 9)

		assert.EqualError(t, err, "Id not found")
	})

	t.Run("delete_ok", func(t *testing.T) {
		repo.
			On("Delete", mock.Anything, mock.AnythingOfType("int")).
			Return(nil).
			Once()

		err := serv.Delete(ctx, 1)

		assert.Equal(t, nil, err)
	})
}
