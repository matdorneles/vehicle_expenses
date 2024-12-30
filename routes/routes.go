package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matdorneles/vehicle_expenses/controllers"
)

func HandleRequests() {
	router := gin.Default()

	// GET Routes
	router.GET("/vehicles", controllers.GetVehicles)
	router.GET("/vehicles/:id", controllers.GetVehicleByID)

	// POST Routes
	router.POST("/vehicles", controllers.PostVehicle)

	// PUT Routes
	router.PATCH("/vehicles/:id", controllers.EditVehicle)

	router.Run(":8080")
}
