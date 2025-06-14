package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medibridge/config"
	"github.com/medibridge/models"
)

type PatientRequest struct {
	FirstName        string `json:"firstName" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Phone           string `json:"phone" binding:"required"`
	DateOfBirth     string `json:"dateOfBirth" binding:"required"`
	Gender          string `json:"gender" binding:"required,oneof=male female other"`
	Address         string `json:"address" binding:"required"`
	EmergencyContact string `json:"emergencyContact" binding:"required"`
	EmergencyPhone  string `json:"emergencyPhone" binding:"required"`
	BloodGroup      string `json:"bloodGroup"`
	Allergies       string `json:"allergies"`
	Diagnosis       string `json:"diagnosis"`
	Notes           string `json:"notes"`
}

type PatientUpdateRequest struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	DateOfBirth     string `json:"dateOfBirth"`
	Gender          string `json:"gender"`
	Address         string `json:"address"`
	EmergencyContact string `json:"emergencyContact"`
	EmergencyPhone  string `json:"emergencyPhone"`
	BloodGroup      string `json:"bloodGroup"`
	Allergies       string `json:"allergies"`
	Diagnosis       string `json:"diagnosis"`
	Notes           string `json:"notes"`
}

func CreatePatient(c *gin.Context) {
	var req PatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse date of birth
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth format. Use YYYY-MM-DD"})
		return
	}

	// Check if email already exists
	var existingPatient models.Patient
	if err := config.DB.Where("email = ?", req.Email).First(&existingPatient).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A patient with this email already exists"})
		return
	}

	userID, _ := c.Get("userID")
	patient := models.Patient{
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Email:           req.Email,
		Phone:           req.Phone,
		DateOfBirth:     dob,
		Gender:          req.Gender,
		Address:         req.Address,
		EmergencyContact: req.EmergencyContact,
		EmergencyPhone:  req.EmergencyPhone,
		BloodGroup:      req.BloodGroup,
		Allergies:       req.Allergies,
		Diagnosis:       req.Diagnosis,
		Notes:           req.Notes,
		CreatedBy:       userID.(uint),
		UpdatedBy:       userID.(uint),
	}

	if err := config.DB.Create(&patient).Error; err != nil {
		// Check for other database errors
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"patients_email_key\" (SQLSTATE 23505)" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "A patient with this email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": patient,
		"message": "Patient created successfully",
	})
}

func GetPatients(c *gin.Context) {
	// Get pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	// Calculate offset
	offset := (page - 1) * limit

	// Build query
	query := config.DB.Model(&models.Patient{})

	// Add search condition if provided
	if search != "" {
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?", 
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count patients"})
		return
	}

	// Get paginated results
	var patients []models.Patient
	if err := query.Offset(offset).Limit(limit).Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	// Calculate total pages
	totalPages := (int(total) + limit - 1) / limit

	// Return paginated response
	c.JSON(http.StatusOK, gin.H{
		"data": patients,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	patientID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var req PatientUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	var patient models.Patient
	if err := config.DB.First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Update fields if provided
	if req.FirstName != "" {
		patient.FirstName = req.FirstName
	}
	if req.LastName != "" {
		patient.LastName = req.LastName
	}
	if req.Email != "" {
		patient.Email = req.Email
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth format. Use YYYY-MM-DD"})
			return
		}
		patient.DateOfBirth = dob
	}
	if req.Gender != "" {
		patient.Gender = req.Gender
	}
	if req.Address != "" {
		patient.Address = req.Address
	}
	if req.EmergencyContact != "" {
		patient.EmergencyContact = req.EmergencyContact
	}
	if req.EmergencyPhone != "" {
		patient.EmergencyPhone = req.EmergencyPhone
	}
	if req.BloodGroup != "" {
		patient.BloodGroup = req.BloodGroup
	}
	if req.Allergies != "" {
		patient.Allergies = req.Allergies
	}

	// Only allow doctors to update diagnosis and notes
	if userRole == models.RoleDoctor {
		if req.Diagnosis != "" {
			patient.Diagnosis = req.Diagnosis
		}
		if req.Notes != "" {
			patient.Notes = req.Notes
		}
	}

	patient.UpdatedBy = userID.(uint)

	if err := config.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": patient,
		"message": "Patient updated successfully",
	})
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	patientID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	if err := config.DB.Delete(&models.Patient{}, patientID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Patient deleted successfully",
	})
} 