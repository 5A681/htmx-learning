package entity

import "time"

type DataTmps struct {
	Id            int       `db:"id"`
	CurrentPhaseA float32   `db:"current_phase_a"`
	CurrentPhaseB float32   `db:"current_phase_b"`
	CurrentPhaseC float32   `db:"current_phase_c"`
	ActivePower   float32   `db:"active_power"`
	ReactivePower float32   `db:"reactive_power"`
	PowerFactor   float32   `db:"power_factor"`
	DataDatetime  time.Time `db:"data_datetime"`
	CreatedAt     time.Time `db:"created_at"`
	BayId         int       `db:"bay_id"`
}
