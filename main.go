package main

import (
	"log"
	"os"

	"pokemonBE/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env hanya jika ada (lokal)
	if os.Getenv("VERCEL") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found (production mode)")
		}
	}
}

func main() {
	godotenv.Load()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
