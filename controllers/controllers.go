package controllers

import (
	"net/http"

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

	if err := c.BindJSON(&newVehicle); err != nil {
		return
	}

	vehicles = append(vehicles, newVehicle)
	c.IndentedJSON(http.StatusCreated, newVehicle)
}

var vehicles = []models.Vehicle{
	{ID: "1", Maker: "Toyota", Model: "Etios", Year: "2014", CreatedAt: "27/12/2024"},
	{ID: "2", Maker: "Fiat", Model: "Toro", Year: "2024", CreatedAt: "27/12/2024"},
	{ID: "3", Maker: "Fiat", Model: "Strada", Year: "2022", CreatedAt: "27/12/2024"},
}
