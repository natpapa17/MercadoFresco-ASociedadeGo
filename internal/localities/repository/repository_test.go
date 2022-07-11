package repository_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/repository"
	"github.com/stretchr/testify/assert"
)



var locality1 = domain.Locality{1, "sp", 1}
var localityReport1 = repository.LocalityReport{1, "Itabaiana", 200}
var localityReport2 = repository.LocalityReport{2, "Nova York", 3000}

func TestCreate(t *testing.T) {
	query := `INSERT INTO locality (name,province_id) VALUES(?,?)`
	t.Run("Create Error saving", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		localityRepo := repository.CreateMySQLRepository(db)
		mock.ExpectPrepare(query).WillReturnError(fmt.Errorf("error"))
		_, err = localityRepo.Create(locality1.Name, locality1.Province_id)
		defer db.Close()
		assert.NotNil(t, err)
	})

	t.Run("Create Ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(locality1.Name, locality1.Province_id).WillReturnResult(sqlmock.NewResult(0, 1))
		localityRepo := repository.CreateMySQLRepository(db)
		_, err = localityRepo.Create( locality1.Name, locality1.Province_id)
		defer db.Close()
		assert.NoError(t, err)
	})
	
}
func TestReportAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	reports := []repository.LocalityReport{localityReport1, localityReport2}
	localityRepo := repository.CreateMySQLRepository(db)

	query := `SELECT locality.id, locality.name, COUNT(seller.locality_id) 
	FROM fresh_market.locality
	INNER JOIN seller ON locality.Id = seller.locality_id
	GROUP BY locality_Id;`
	t.Run("ReportAll - OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"LocalityId",
			"LocalityName",
			"SellersCount",
		}).AddRow(
			reports[0].Locality_id,
			reports[0].Name,
			reports[0].SellersCount,
		).AddRow(
			reports[1].Locality_id,
			reports[1].Name,
			reports[1].SellersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		result, err := localityRepo.ReportAll()
		assert.NoError(t, err)

		assert.Equal(t, result[0].Locality_id, reports[0].Locality_id)
		assert.Equal(t, result[1].Locality_id, reports[1].Locality_id)
	})
	t.Run("GenerateReportAll - Fail Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"locality_id",
			"Name",
			"SellersCount",
		}).AddRow("", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		_, err = localityRepo.ReportAll()
		assert.Error(t, err)
	})
	t.Run("GetAll - Fail Select/Read", func(t *testing.T) {
		_, err = localityRepo.ReportAll()
		assert.Error(t, err)
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
	})
}

func TestReportById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	localityReports := []repository.LocalityReport{localityReport1, localityReport2}
	localityRepo := repository.CreateMySQLRepository(db)

	rows := sqlmock.NewRows([]string{
		"Locality_id",
		"Name",
		"SellersCount",
	}).AddRow(
		localityReports[0].Locality_id,
		localityReports[0].Name,
		localityReports[0].SellersCount,
	)

	query := `SELECT locality.id, locality.name, COUNT(seller.locality_id) 
	FROM fresh_market.locality
	INNER JOIN seller ON locality.id = seller.locality_id
	WHERE locality.id = ?
	GROUP BY locality_id;`
	t.Run(" report by ID - OK", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)
		result, _ := localityRepo.ReportById(1)
		assert.NoError(t, err)

		assert.Equal(t, localityReports[0].Locality_id, result.Locality_id)
	})
	t.Run("Generate report by ID - Fail prepar query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(fmt.Errorf("Fail to prepar query"))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Fail to prepar query"))
		localityRepo := repository.CreateMySQLRepository(db)
		_, err = localityRepo.ReportById(1)
		assert.Equal(t, fmt.Errorf("Fail to prepar query"), err)

	})
	t.Run("Get ID - Locality Not found", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Locality 1 not found"))
		localityRepo := repository.CreateMySQLRepository(db)
		_, err = localityRepo.ReportById(1)
		assert.Equal(t, fmt.Errorf("Locality 1 not found"), err)

	})
}

