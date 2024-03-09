package domain

import "time"

type Todo struct {
	ID              int        `json:"id"`
	ActivityGroupID int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        bool       `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
