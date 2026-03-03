package handlers

import "github.com/gin-gonic/gin"

type responseMeta struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func respondSuccess(c *gin.Context, code int, message string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"code":    code,
		"data":    data,
	})
}

func respondError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"code":    code,
	})
}
