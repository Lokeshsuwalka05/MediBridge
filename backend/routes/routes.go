package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/medibridge/controllers"
	"github.com/medibridge/middleware"
	"github.com/medibridge/models"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	// Receptionist routes
	receptionist := authorized.Group("/receptionist")
	receptionist.Use(middleware.RoleMiddleware(models.RoleReceptionist))
	{
		receptionist.POST("/patients", controllers.CreatePatient)
		receptionist.GET("/patients", controllers.GetPatients)
		receptionist.PUT("/patients/:id", controllers.UpdatePatient)
		receptionist.DELETE("/patients/:id", controllers.DeletePatient)
	}

	// Doctor routes
	doctor := authorized.Group("/doctor")
	doctor.Use(middleware.RoleMiddleware(models.RoleDoctor))
	{
		doctor.GET("/patients", controllers.GetPatients)
		doctor.PATCH("/patients/:id", controllers.UpdatePatient)
	}
} 