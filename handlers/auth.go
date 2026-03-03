package handlers

import (
	"net/http"

	"pokemonBE/config"
	"pokemonBE/services"

	"github.com/gin-gonic/gin"
)

// payload structs

type registerInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register creates a new user
func Register(c *gin.Context) {
	var input registerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondError(c, http.StatusBadRequest, "input tidak valid")
		return
	}

	user, err := services.RegisterUser(input.Username, input.Password)
	if err != nil {
		if err.Error() == "username already taken" {
			respondError(c, http.StatusBadRequest, "username sudah digunakan")
		} else {
			respondError(c, http.StatusInternalServerError, "gagal register")
		}
		return
	}

	respondSuccess(c, http.StatusCreated, "register berhasil", gin.H{"user": user})
}

// Login authenticates an existing user
func Login(c *gin.Context) {
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondError(c, http.StatusBadRequest, "input tidak valid")
		return
	}

	user, err := services.AuthenticateUser(input.Username, input.Password)
	if err != nil {
		respondError(c, http.StatusUnauthorized, "username atau password salah")
		return
	}

	token, _ := config.GenerateJWT(user.ID)

	respondSuccess(c, http.StatusOK, "login berhasil", gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"saldo_uang": user.SaldoUang,
		},
	})
}
