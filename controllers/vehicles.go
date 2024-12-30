package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matdorneles/vehicle_expenses/models"
)

func GetVehicles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vehicles)
}

func GetVehicleByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range vehicles {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vehicle not found"})
}

func PostVehicle(c *gin.Context) {
	var newVehicle models.Vehicle
	var now time.Time = time.Now()

	if err := c.ShouldBindJSON(&newVehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newVehicle.CreatedAt = now.Format("01-02-2006")
	vehicles = append(vehicles, newVehicle)
	c.IndentedJSON(http.StatusCreated, newVehicle)
}

func EditVehicle(c *gin.Context) {
	var editVehicle models.Vehicle
	id := c.Param("id")

	if err := c.ShouldBindJSON(&editVehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, v := range vehicles {
		if v.ID == id {
			v.Maker = editVehicle.Maker
			v.Model = editVehicle.Model
			v.Year = editVehicle.Year
			c.IndentedJSON(http.StatusAccepted, v)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Not Found"})
}

// Temporary until database is implemented
var vehicles = []models.Vehicle{
	{ID: "1", Maker: "Toyota", Model: "Etios", Year: "2014", CreatedAt: "27/12/2024"},
	{ID: "2", Maker: "Fiat", Model: "Toro", Year: "2024", CreatedAt: "27/12/2024"},
	{ID: "3", Maker: "Fiat", Model: "Strada", Year: "2022", CreatedAt: "27/12/2024"},
}
