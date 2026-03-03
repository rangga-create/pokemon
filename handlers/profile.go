package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile returns data for the authenticated user.
func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	respondSuccess(c, http.StatusOK, "ini endpoint yang diproteksi JWT", gin.H{
		"user_id": userID,
	})
}
