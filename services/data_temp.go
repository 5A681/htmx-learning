package services

import (
	"htmx-learning/entity"
	"htmx-learning/filter"
	"log"
	"time"
)

func (s service) GetLatestData(bayId int, filter filter.SortData) ([]entity.DataTmps, error) {

	maxdate, err := s.repo.GetMaxDate()
	if err != nil {
		return nil, err
	}

	if maxdate != nil {
		data, err := s.repo.GetLatestDataByBayId(bayId, filter, *maxdate)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, nil

}

func (s service) GetDataLatestMonthDayTime(bayId int, filter filter.SortData) ([]entity.DataTmps, error) {
	var datas []entity.DataTmps
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 8, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 15, 30, 0, 0, time.Local)
			data, err := s.repo.GetMaxDataPerDayPerTime(bayId, min, max)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				datas = append(datas, *data)
			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 8, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 15, 30, 0, 0, time.Local)
			data, err := s.repo.GetMaxDataPerDayPerTime(bayId, min, max)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				datas = append(datas, *data)
			}
		}
	}

	return datas, nil

}

func (s service) GetDataLatestMonthNightTime(bayId int, filter filter.SortData) ([]entity.DataTmps, error) {
	var datas []entity.DataTmps
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 7, 30, 0, 0, time.Local)
			min2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 16, 0, 0, 0, time.Local)
			max2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
			data, err := s.repo.GetMaxDataPerDayPerTimeTwoTime(bayId, min, max, min2, max2)
			if err != nil {
				log.Println(err.Error(), i)
			}
			if data != nil {
				datas = append(datas, *data)
			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 7, 30, 0, 0, time.Local)
			min2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 16, 0, 0, 0, time.Local)
			max2 := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
			data, err := s.repo.GetMaxDataPerDayPerTimeTwoTime(bayId, min, max, min2, max2)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				datas = append(datas, *data)
			}
		}
	}

	return datas, nil

}
func (s service) GetDataLatestMonthAllTime(bayId int, filter filter.SortData) ([]entity.DataTmps, error) {
	var datas []entity.DataTmps
	err := s.repo.CheckPreviousMonth()
	if err != nil {
		log.Println(err.Error())
		currentTime, err := s.repo.GetMaxDate()
		if err != nil {
			return nil, err
		}
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
			data, err := s.repo.GetMinDataPerDayPerTime(bayId, min, max)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				datas = append(datas, *data)
			}
		}

	} else {
		currentTime := time.Now()
		currentTime = currentTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			min := time.Date(currentTime.Year(), currentTime.Month(), i+1, 0, 0, 0, 0, time.Local)
			max := time.Date(currentTime.Year(), currentTime.Month(), i+1, 23, 30, 0, 0, time.Local)
			data, err := s.repo.GetMinDataPerDayPerTime(bayId, min, max)
			if err != nil {
				log.Println(err.Error())
			}
			if data != nil {
				datas = append(datas, *data)
			}
		}
	}

	return datas, nil

}

func (s service) GetDataLatestYearPeakTime(bayId int, year int, filter filter.SortData) ([]entity.DataTmps, error) {
	var datas []entity.DataTmps

	for i := 0; i < 12; i++ {

		data, err := s.repo.GetMaxDataPerMonth(bayId, year, i+1)
		if err != nil {
			log.Println(err.Error(), 1)
		}
		if data != nil {
			datas = append(datas, *data)
		}
	}

	return datas, nil

}

func (s service) GetDataLatestYearLightTime(bayId int, year int, filter filter.SortData) ([]entity.DataTmps, error) {

	var datas []entity.DataTmps

	for i := 0; i < 12; i++ {

		data, err := s.repo.GetMinDataPerMonth(bayId, year, i+1)
		if err != nil {
			log.Println(err.Error(), 3)
		}
		if data != nil {
			datas = append(datas, *data)
		}
	}

	return datas, nil

}
