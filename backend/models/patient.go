package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	FirstName        string         `gorm:"not null" json:"firstName"`
	LastName         string         `gorm:"not null" json:"lastName"`
	Email           string         `gorm:"not null;unique" json:"email"`
	Phone           string         `gorm:"not null" json:"phone"`
	DateOfBirth     time.Time      `gorm:"not null" json:"dateOfBirth"`
	Gender          string         `gorm:"not null" json:"gender"`
	Address         string         `gorm:"not null" json:"address"`
	EmergencyContact string         `gorm:"not null" json:"emergencyContact"`
	EmergencyPhone  string         `gorm:"not null" json:"emergencyPhone"`
	BloodGroup      string         `json:"bloodGroup"`
	Allergies       string         `json:"allergies"`
	Diagnosis       string         `json:"diagnosis"`
	Notes           string         `json:"notes"`
	CreatedBy       uint           `gorm:"not null" json:"createdBy"`
	UpdatedBy       uint           `gorm:"not null" json:"updatedBy"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
} 