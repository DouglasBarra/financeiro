package models

import "time"

type Category struct {
	ID          uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
