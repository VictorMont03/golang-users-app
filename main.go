package main

import (
	"log"

	"github.com/VictorMont03/golang-users-app/src/config/database/mongodb"
	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection

	database, err := mongodb.NewMongoDBConnection()

	if err != nil {
		log.Fatalf("Error connecting to database %s", err.Error())
		return
	}

	// Init dependencies
	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server")
	}
}
