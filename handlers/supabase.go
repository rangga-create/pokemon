package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetSupabaseConfig returns Supabase URL and public API key from environment variables.
func GetSupabaseConfig(c *gin.Context) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_API_KEY")

	if url == "" || key == "" {
		respondError(c, http.StatusInternalServerError, "SUPABASE_URL atau SUPABASE_API_KEY belum diset")
		return
	}

	respondSuccess(c, http.StatusOK, "supabase config", gin.H{
		"supabase_url":     url,
		"supabase_api_key": key,
	})
}
