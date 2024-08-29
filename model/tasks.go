package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title         string `gorm:"type:varchar(100)"`
	Description   string `gorm:"type:text"`
	CreatedByUser uint
}
