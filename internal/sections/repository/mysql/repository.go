package mysql

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
)

type mySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) sections.Repository {
	return mySQLRepository{
		db: db,
	}
}

func (m mySQLRepository) GetAll(ctx context.Context) ([]sections.Section, error) {
	sects := []sections.Section{}

	rows, err := m.db.QueryContext(ctx, "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section")
	if err != nil {
		return sects, err
	}

	defer rows.Close()

	for rows.Next() {
		var sect sections.Section

		err := rows.Scan(&sect.ID, &sect.SectionNumber, &sect.CurrentTemperature, &sect.MinimumTemperature, &sect.CurrentCapacity, &sect.MinimumCapacity, &sect.MaximumCapacity, &sect.WarehouseID, &sect.ProductTypeID)
		if err != nil {
			return sects, err
		}

		sects = append(sects, sect)
	}

	return sects, nil
}

func (m mySQLRepository) GetById(ctx context.Context, id int) (sections.Section, error) {
	var sect sections.Section

	row := m.db.QueryRowContext(
		ctx,
		"SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM section WHERE id=?",
		id,
	)

	err := row.Scan(&sect.ID, &sect.SectionNumber, &sect.CurrentTemperature, &sect.MinimumTemperature, &sect.CurrentCapacity, &sect.MinimumCapacity, &sect.MaximumCapacity, &sect.WarehouseID, &sect.ProductTypeID)
	if err != nil {
		return sections.Section{}, err
	}

	return sect, nil
}

func (m mySQLRepository) LastID(ctx context.Context) (int, error) {
	var maxCount sql.NullInt64

	row := m.db.QueryRowContext(ctx, "SELECT MAX(id) as last_id FROM section")

	err := row.Scan(&maxCount)
	if err != nil {
		return 0, nil
	}

	if maxCount.Valid {
		return int(maxCount.Int64), nil
	}

	return 0, nil
}

func (m mySQLRepository) HasSectionNumber(ctx context.Context, number int) (bool, error) {
	var sect_num sql.NullInt64

	row := m.db.QueryRowContext(
		ctx,
		"SELECT section_number FROM section WHERE section_number=?",
		number,
	)

	err := row.Scan(&sect_num)
	if err != nil {
		return false, nil
	}

	return sect_num.Valid, nil
}

func (m mySQLRepository) Add(ctx context.Context, id int, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (sections.Section, error) {
	sect := sections.Section{
		ID:                 id,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemprarature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}

	_, err := m.db.ExecContext(
		ctx,
		"INSERT INTO section (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		&sect.ID,
		&sect.SectionNumber,
		&sect.CurrentTemperature,
		&sect.MinimumTemperature,
		&sect.CurrentCapacity,
		&sect.MinimumCapacity,
		&sect.MaximumCapacity,
		&sect.WarehouseID,
		&sect.ProductTypeID,
	)

	if err != nil {
		return sect, err
	}

	return sect, nil
}

func (m mySQLRepository) UpdateById(ctx context.Context, id int, section sections.Section) (sections.Section, error) {

	sect, err := m.GetById(ctx, id)
	if err != nil {
		return sections.Section{}, err
	}

	sectTypes := reflect.TypeOf(sect)

	sectionValues := reflect.ValueOf(&section).Elem()
	sectValues := reflect.ValueOf(&sect).Elem()

	for i := sectTypes.NumField() - 1; i > 0; i-- {
		if !sectionValues.Field(i).IsZero() {
			sectValues.Field(i).Set(sectionValues.Field(i))
		}
	}

	_, err = m.db.ExecContext(
		ctx,
		"UPDATE section SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?",
		&sect.SectionNumber,
		&sect.CurrentTemperature,
		&sect.MinimumTemperature,
		&sect.CurrentCapacity,
		&sect.MinimumCapacity,
		&sect.MaximumCapacity,
		&sect.WarehouseID,
		&sect.ProductTypeID,
		&sect.ID,
	)

	if err != nil {
		return sections.Section{}, err
	}

	return sect, nil
}

func (m mySQLRepository) Delete(ctx context.Context, id int) error {
	_, err := m.GetById(ctx, id)
	if err != nil {
		return err
	}

	rows, err := m.db.ExecContext(ctx, "DELETE FROM section WHERE id=?", id)
	if err != nil {
		return err
	}

	if affected, _ := rows.RowsAffected(); affected == 0 {
		log.Println("Delete", affected)
		return err
	}

	return nil
}
