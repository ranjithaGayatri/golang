package main

import (
	"PackageApi/models"
	"PackageApi/router"
	"fmt"
)

func main() {
	fmt.Println("Starting in Main")
	models.ConnectingToDatabase()
	router.RoutesDefined()
	fmt.Println("Done with main...")
}
