package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/medibridge/config"
	"github.com/medibridge/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})

	// Create test user
	password, _ := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
	user := models.User{
		Name:         "Test User",
		Email:        "test@example.com",
		PasswordHash: string(password),
		Role:         models.RoleDoctor,
	}
	config.DB.Create(&user)

	// Test cases
	tests := []struct {
		name       string
		payload    LoginRequest
		wantStatus int
		wantError  bool
	}{
		{
			name: "Valid credentials",
			payload: LoginRequest{
				Email:    "test@example.com",
				Password: "test123",
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name: "Invalid password",
			payload: LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpass",
			},
			wantStatus: http.StatusUnauthorized,
			wantError:  true,
		},
		{
			name: "Invalid email",
			payload: LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "test123",
			},
			wantStatus: http.StatusUnauthorized,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			body, _ := json.Marshal(tt.payload)
			c.Request = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))

			// Call handler
			Login(c)

			// Assertions
			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantError {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "error")
			} else {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "token")
				assert.Contains(t, response, "user")
			}
		})
	}
} 