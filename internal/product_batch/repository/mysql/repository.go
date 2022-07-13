package mysql

import (
	"context"
	"database/sql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
)

type mySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) product_batch.Repository {
	return mySQLRepository{
		db: db,
	}
}

func (m mySQLRepository) GetById(ctx context.Context, id int) ([]product_batch.ProductsReport, error) {
	var prl []product_batch.ProductsReport

	rows, err := m.db.QueryContext(
		ctx,
		"SELECT product_batch.section_id, section.section_number, product_batch.current_quantity as products_count FROM product_batch JOIN section ON product_batch.section_id=section.id WHERE product_batch.section_id=?",
		id,
	)
	defer rows.Close()

	if err != nil {
		return []product_batch.ProductsReport{}, err
	}

	for rows.Next() {
		var pr product_batch.ProductsReport

		err := rows.Scan(&pr.SectionID, &pr.SectionNumber, &pr.ProductsCount)
		if err != nil {
			return []product_batch.ProductsReport{}, err
		}

		prl = append(prl, pr)

	}

	return prl, nil
}

func (m mySQLRepository) Add(ctx context.Context, id int, batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature int, productID int, sectionID int) (product_batch.ProductBatch, error) {
	pb := product_batch.ProductBatch{
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

	_, err := m.db.ExecContext(
		ctx,
		"INSERT INTO product_batch (id, batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		id,
		batchNumber,
		currentQuantity,
		currentTemperature,
		dueDate,
		initialQuantity,
		manufacturingDate,
		manufacturingHour,
		minimumTemperature,
		productID,
		sectionID,
	)

	if err != nil {
		return product_batch.ProductBatch{}, err
	}

	return pb, nil
}

func (m mySQLRepository) LastID(ctx context.Context) (int, error) {
	var maxCount sql.NullInt64

	row := m.db.QueryRowContext(ctx, "SELECT MAX(id) as last_id FROM product_batch")

	err := row.Scan(&maxCount)
	if err != nil {
		return 0, nil
	}

	if maxCount.Valid {
		return int(maxCount.Int64), nil
	}

	return 0, nil
}

func (m mySQLRepository) HasBatchNumber(ctx context.Context, number int) (bool, error) {
	var batchNumber sql.NullInt64

	row := m.db.QueryRowContext(
		ctx,
		"SELECT batch_number FROM product_batch WHERE batch_number=?",
		number,
	)

	err := row.Scan(&batchNumber)
	if err != nil {
		return false, nil
	}

	return batchNumber.Valid, nil
}
