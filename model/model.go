package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	EGN    string `json:"egn"`
	Status bool   `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}
func (e *Employee) Enable() {
	e.Status = true
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
