package model

import "time"

type RatingScore struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	Score     int       `db:"score"`
}

type CategoryScore struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	Score     int       `db:"score"`
}

type TicketScore struct {
	TicketID  int       `db:"ticket_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	Score     int       `db:"score"`
}
