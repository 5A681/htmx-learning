package repository

import "htmx-learning/entity"

func (s repository) GetSubStationById(id int) (*entity.SubStation, error) {
	var subStation entity.SubStation
	err := s.db.Get(&subStation, `select * from sub_stations where id = $1 order by id asc`, id)
	if err != nil {
		return nil, err
	}
	return &subStation, nil
}
