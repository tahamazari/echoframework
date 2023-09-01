package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
