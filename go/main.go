package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamjaekim/random-restaurant-api/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.Use(cors.Default())

	// router.GET("/", handlers.Index)
	router.GET("/api/stores/zip/:zipCode", handlers.GetRestaurants)
	router.GET("/api/stores/single/:storeId", handlers.GetRestaurant)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}