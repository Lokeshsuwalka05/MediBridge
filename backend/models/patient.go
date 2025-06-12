package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Age       int            `gorm:"not null" json:"age"`
	Gender    string         `gorm:"not null" json:"gender"`
	Symptoms  string         `json:"symptoms"`
	Diagnosis string         `json:"diagnosis"`
	Notes     string         `json:"notes"`
	CreatedBy uint           `gorm:"not null" json:"created_by"`
	UpdatedBy uint           `gorm:"not null" json:"updated_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
} 