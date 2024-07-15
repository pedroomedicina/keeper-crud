package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model        // Adds fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"type:varchar(255)"`
}
