package main

import (
	"github.com/jinzhu/gorm"
)

type ModelExample struct {
	gorm.Model
	ID int64 `gorm:"auto_increment"`
}
