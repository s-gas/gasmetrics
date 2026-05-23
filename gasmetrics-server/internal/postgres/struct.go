package postgres

import "time"

type Reading struct {
	Id		int				`json:"id"`
	Value int 			`json:"value"`
	Date 	time.Time `json:"recorded_at"`
}
