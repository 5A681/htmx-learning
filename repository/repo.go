package repository

import (
	"htmx-learning/entity"
	"htmx-learning/filter"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetSubStationById(id int) (*entity.SubStation, error)
	GetBayById(id int) (*entity.Bay, error)
	GetDataTmpsById(id int) (*entity.DataTmps, error)
	CreateDataTmep(data entity.DataTmps) error
	GetDataTempsByBayId(bayId int, sort filter.SortData) ([]entity.DataTmps, error)
	GetLatestDataByBayId(bayId int, sort filter.SortData, date time.Time) ([]entity.DataTmps, error)
	GetMaxDate() (*time.Time, error)
	GetMaxDataPerDayPerTime(bayId int, minTime time.Time, maxTime time.Time) (*entity.DataTmps, error)
	CheckPreviousMonth() error
	CheckPreviousYear() error
	GetMaxDataPerDayPerTimeTwoTime(bayId int, minTime1 time.Time, maxTime1 time.Time, minTime2 time.Time, maxTime2 time.Time) (*entity.DataTmps, error)
	GetMinDataPerDayPerTime(bayId int, minTime time.Time, maxTime time.Time) (*entity.DataTmps, error)
	GetMaxDataPerMonth(bayId int, year int, month int) (*entity.DataTmps, error)
	GetMinDataPerMonth(bayId int, year int, month int) (*entity.DataTmps, error)
	GetAllYears() ([]int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return repository{db}
}

func (r repository) GetMaxDate() (*time.Time, error) {
	var maxDateMinusOne time.Time

	// Execute the query
	query := `
        SELECT MAX(dt.data_datetime) - INTERVAL '1 day' AS max_date_minus_one
        FROM data_tmps dt limit 1;
    `
	err := r.db.Get(&maxDateMinusOne, query)
	if err != nil {
		return nil, err
	}
	return &maxDateMinusOne, nil
}
func (r repository) CheckPreviousMonth() error {
	var maxDateMinusOne time.Time

	// Execute the query
	currentTime := time.Now()
	query := `
        SELECT * FROM data_tmps dt where EXTRACT(MONTH FROM dt.data_datetime) = $1;
    `
	err := r.db.Get(&maxDateMinusOne, query, currentTime.Month()-1)
	if err != nil {
		return err
	}
	log.Println("maxtime ", maxDateMinusOne)
	return nil
}

func (r repository) CheckPreviousYear() error {
	var maxDateMinusOne time.Time

	// Execute the query
	currentTime := time.Now()
	query := `
        SELECT data_datetime FROM data_tmps dt where EXTRACT(YEAR FROM dt.data_datetime) = $1 limit 1;
    `
	err := r.db.Get(&maxDateMinusOne, query, currentTime.Year()-1)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) GetAllYears() ([]int, error) {
	var years []int
	err := r.db.Select(&years, `SELECT DISTINCT EXTRACT(YEAR FROM data_datetime) AS year
				FROM data_tmps ORDER BY year desc;`)
	if err != nil {
		return nil, err
	}
	return years, nil
}
