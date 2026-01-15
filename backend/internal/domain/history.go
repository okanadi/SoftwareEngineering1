package domain

import "time"

type HistoryDB struct {
	ID        string    `db:"id"`
	StepId    string    `db:"step_id"`
	Status    string    `db:"status"`
	Note      string    `db:"note"`
	UserName  string    `db:"user_name"`
	Timestamp time.Time `db:"timestamp"`
	Photos    []string  `db:"photos"`
}
