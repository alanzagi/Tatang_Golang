package entity

import "time"

type Debt struct {
	ID        int64
	Name      string
	Amount    int64
	Status    string
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
