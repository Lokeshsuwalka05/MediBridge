package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medibridge/config"
	"github.com/medibridge/models"
	"github.com/medibridge/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Login attempt for email: %s", req.Email)
	log.Printf("Request body: %+v", req)

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		log.Printf("User not found: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	log.Printf("User found: %+v", user)
	log.Printf("Stored password hash: %s", user.PasswordHash)
	log.Printf("Attempting to compare with provided password: %s", req.Password)

	// Log the database connection status
	if err := config.DB.Raw("SELECT 1").Error; err != nil {
		log.Printf("Database connection error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	// Log all users in the database for debugging
	var allUsers []models.User
	if err := config.DB.Find(&allUsers).Error; err != nil {
		log.Printf("Error fetching all users: %v", err)
	} else {
		log.Printf("All users in database: %+v", allUsers)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		log.Printf("Password mismatch: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	log.Printf("Password verified successfully")

	token, err := utils.GenerateToken(&user)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	log.Printf("Login successful for user: %s", user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
} 