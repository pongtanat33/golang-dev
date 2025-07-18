package entities

import "time"

type ActivityFrequency struct {
	Id           string     `db:"id"`
	ActivityId       string     `db:"activity_id"`
	Type         string     `db:"type"`
	SelectedDate *time.Time `db:"selected_date"`
	CreatedAt    time.Time  `db:"created_at"`
	CreatedBy    string     `db:"created_by"`
	UpdatedAt    *time.Time `db:"updated_at"`
	UpdatedBy    string     `db:"updated_by"`
}
