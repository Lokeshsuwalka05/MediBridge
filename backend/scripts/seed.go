package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/medibridge/config"
	"github.com/medibridge/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database
	config.InitDB()

	// Auto migrate the schema
	config.DB.AutoMigrate(&models.User{})

	// Create doctor
	doctorPassword, _ := bcrypt.GenerateFromPassword([]byte("doctor@#123"), bcrypt.DefaultCost)
	doctor := models.User{
		Name:         "Dr. John Doe",
		Email:        "doctor@medibridge.com",
		PasswordHash: string(doctorPassword),
		Role:         models.RoleDoctor,
	}

	// Create receptionist
	receptionistPassword, _ := bcrypt.GenerateFromPassword([]byte("reception@#123"), bcrypt.DefaultCost)
	receptionist := models.User{
		Name:         "Jane Smith",
		Email:        "receptionist@medibridge.com",
		PasswordHash: string(receptionistPassword),
		Role:         models.RoleReceptionist,
	}

	// Check if users already exist
	var existingDoctor models.User
	if err := config.DB.Where("email = ?", doctor.Email).First(&existingDoctor).Error; err != nil {
		if err := config.DB.Create(&doctor).Error; err != nil {
			log.Printf("Error creating doctor: %v", err)
		} else {
			log.Println("Doctor user created successfully")
		}
	} else {
		log.Println("Doctor user already exists")
	}

	var existingReceptionist models.User
	if err := config.DB.Where("email = ?", receptionist.Email).First(&existingReceptionist).Error; err != nil {
		if err := config.DB.Create(&receptionist).Error; err != nil {
			log.Printf("Error creating receptionist: %v", err)
		} else {
			log.Println("Receptionist user created successfully")
		}
	} else {
		log.Println("Receptionist user already exists")
	}
} 