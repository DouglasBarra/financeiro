package models

import "time"

type Card struct {
	ID        uint
	Number    int64
	BankID    int64
	Bank      Bank
	CreatedAt time.Time
	UpdatedAt time.Time
}
