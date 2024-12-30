package main

import (
	"github.com/matdorneles/vehicle_expenses/database"
	"github.com/matdorneles/vehicle_expenses/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
