package repository

import "htmx-learning/entity"

func (s repository) GetBayById(id int) (*entity.Bay, error) {
	var bay entity.Bay
	err := s.db.Get(&bay, `select * from bays where id = $1 order by id asc`, id)
	if err != nil {
		return nil, err
	}
	return &bay, nil
}
