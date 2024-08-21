package entity

import "time"

type Bay struct {
	Id           int       `db:"id"`
	Name         string    `db:"name"`
	CreatedAt    time.Time `db:"created_at"`
	SubStationId int       `db:"sub_station_id"`
}
