package mysql_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections/repository/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSectionWithId(id int, sectionNumber int, currentTemperature float32, minimumTemperature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) sections.Section {
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}

	t.Run("get_all_ok", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1, 2, 1, 4, 1, 1).AddRow(2, 2, 1, 1, 5, 2, 10, 1, 1))

		sects, err := repo.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, []sections.Section{{1, 1, 1, 1, 2, 1, 4, 1, 1}, {2, 2, 1, 1, 5, 2, 10, 1, 1}}, sects)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("get_all_fail", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section").
			WillReturnError(errors.New("error"))

		_, err := repo.GetAll(context.Background())
		assert.Error(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}

	t.Run("get_by_id_ok", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1, 2, 1, 4, 1, 1))

		sect, err := repo.GetById(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, sections.Section{1, 1, 1, 1, 2, 1, 4, 1, 1}, sect)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("get_by_id_fail", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?").
			WillReturnError(errors.New("error"))

		_, err := repo.GetById(context.Background(), 1)
		assert.Error(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestLastID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	columns := []string{"last_id"}

	t.Run("last_id_ok", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT MAX\\(id\\) as last_id FROM section").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1))

		id, err := repo.LastID(context.Background())

		assert.Equal(t, 1, id)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("last_id_fail", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT MAX\\(id\\) as last_id FROM section").
			WillReturnError(errors.New("error"))

		_, err := repo.LastID(context.Background())

		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestHasSectionNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	columns := []string{"section_number"}

	t.Run("has_section_number_ok", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT section_number FROM section WHERE section_number=?").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1))

		has, err := repo.HasSectionNumber(context.Background(), 1)
		assert.Equal(t, true, has)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)

	})

	t.Run("has_section_number_fail", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT section_number FROM section WHERE section_number=?").
			WillReturnError(errors.New("error"))

		_, err := repo.HasSectionNumber(context.Background(), 1)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	t.Run("add_ok", func(t *testing.T) {
		mock.
			ExpectExec("INSERT INTO section").
			WithArgs(1, 1, 1.0, 1.0, 2, 1, 4, 1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sect, err := repo.Add(context.Background(), 1, 1, 1.0, 1.0, 2, 1, 4, 1, 1)
		assert.Equal(t, sections.Section{1, 1, 1, 1, 2, 1, 4, 1, 1}, sect)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("add_fail", func(t *testing.T) {
		mock.
			ExpectExec("INSERT INTO section").
			WithArgs(1, 1, 1.0, 1.0, 2, 1, 4, 1, 1).
			WillReturnError(errors.New("error"))

		_, err := repo.Add(context.Background(), 1, 1, 1.0, 1.0, 2, 1, 4, 1, 1)
		assert.Error(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}

	t.Run("update_by_id_ok", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1, 2, 1, 4, 1, 1))

		mock.
			ExpectExec("UPDATE section").
			WithArgs(1, 1.0, 1.0, 5, 1, 10, 1, 1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sect, err := repo.UpdateById(context.Background(), 1, sections.Section{0, 0, 0.0, 0.0, 5, 1, 10, 0, 0})
		assert.Equal(t, sections.Section{1, 1, 1.0, 1.0, 5, 1, 10, 1, 1}, sect)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("update_by_id_fail", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?").
			WillReturnError(errors.New("error"))

		_, err := repo.UpdateById(context.Background(), 1, sections.Section{0, 0, 0.0, 0.0, 5, 1, 10, 0, 0})
		assert.Error(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mysql.NewMySQLRepository(db)

	columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}

	t.Run("delete_ok", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1, 2, 1, 4, 1, 1))

		mock.
			ExpectExec("DELETE FROM section WHERE id=?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(context.Background(), 1)
		assert.Nil(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})

	t.Run("delete_fail", func(t *testing.T) {
		mock.
			ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1, 1, 1, 2, 1, 4, 1, 1))

		err := repo.Delete(context.Background(), 1)
		assert.Error(t, err)
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}
