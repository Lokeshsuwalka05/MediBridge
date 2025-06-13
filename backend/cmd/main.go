package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/medibridge/config"
	"github.com/medibridge/models"
	"github.com/medibridge/routes"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database
	log.Println("Initializing database connection...")
	config.InitDB()
	log.Println("Database connection established")

	// Auto migrate the schema
	log.Println("Running database migrations...")
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})
	log.Println("Database migrations completed")

	// Seed initial users if they don't exist
	log.Println("Checking for default users...")
	seedUsers()

	// Initialize Gin router
	r := gin.Default()

	// Add logging middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Configure CORS
	log.Println("Configuring CORS...")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // In production, replace with your frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	log.Println("Setting up routes...")
	routes.SetupRoutes(r)
	log.Println("Routes setup completed")

	// Start server
	port := os.Getenv("PORT") // Render uses PORT environment variable
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func seedUsers() {
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