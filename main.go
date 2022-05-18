package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-restaurent-management/database"
	"os"
	//"golang-restaurent-management/controllers"
	"golang-restaurent-management/routes"
	//"golang-restaurent-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	//"github.com/labstack/echo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	//router.Use(middleware.Authentication())
	routes.MenuRoutes(router)

	fmt.Println("server running on port 8080")
	fmt.Println("testing changes feom sublime merge")
	router.Run(":8080")

}
