package models

import "time"

type Vehicle struct {
	VehicleID   uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Location    string
	Price       int
	Description string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
