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
	config.InitDB()

	// Auto migrate the schema
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})

	// Seed initial users if they don't exist
	seedUsers()

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func seedUsers() {
	// Check if users exist
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	// Create doctor
	doctorPassword, _ := bcrypt.GenerateFromPassword([]byte("doctor@#123"), bcrypt.DefaultCost)
	doctor := models.User{
		Name:         "Dr. John Doe",
		Email:        "doctor@medibridge.com",
		PasswordHash: string(doctorPassword),
		Role:         models.RoleDoctor,
	}
	config.DB.Create(&doctor)

	// Create receptionist
	receptionistPassword, _ := bcrypt.GenerateFromPassword([]byte("reception@#123"), bcrypt.DefaultCost)
	receptionist := models.User{
		Name:         "Jane Smith",
		Email:        "receptionist@medibridge.com",
		PasswordHash: string(receptionistPassword),
		Role:         models.RoleReceptionist,
	}
	config.DB.Create(&receptionist)
} 