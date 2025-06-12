package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medibridge/config"
	"github.com/medibridge/models"
)

type PatientRequest struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Symptoms string `json:"symptoms"`
}

type PatientUpdateRequest struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Symptoms  string `json:"symptoms"`
	Diagnosis string `json:"diagnosis"`
	Notes     string `json:"notes"`
}

func CreatePatient(c *gin.Context) {
	var req PatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	patient := models.Patient{
		Name:      req.Name,
		Age:       req.Age,
		Gender:    req.Gender,
		Symptoms:  req.Symptoms,
		CreatedBy: userID.(uint),
		UpdatedBy: userID.(uint),
	}

	if err := config.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
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

	// Update fields based on role
	if userRole == models.RoleDoctor {
		if req.Diagnosis != "" {
			patient.Diagnosis = req.Diagnosis
		}
		if req.Notes != "" {
			patient.Notes = req.Notes
		}
	} else {
		if req.Name != "" {
			patient.Name = req.Name
		}
		if req.Age != 0 {
			patient.Age = req.Age
		}
		if req.Gender != "" {
			patient.Gender = req.Gender
		}
		if req.Symptoms != "" {
			patient.Symptoms = req.Symptoms
		}
	}

	patient.UpdatedBy = userID.(uint)

	if err := config.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
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

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
} 