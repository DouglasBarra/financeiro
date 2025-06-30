package models

import "time"

type Bank struct {
	ID          uint
	Name        string
	BankCode    int64
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
