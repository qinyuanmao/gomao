package db

import "time"

type Model struct {
	ID        uint64 `gorm:"primary_key" type:bigint unsigned"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
