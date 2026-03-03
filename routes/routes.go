package routes

import (
	"pokemonBE/handlers"
	"pokemonBE/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/supabase", handlers.GetSupabaseConfig)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", handlers.Profile)
}
