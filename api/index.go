package handler

import (
    "net/http"

    "pokemonBE/config"
    "pokemonBE/routes"

    "github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
    config.ConnectDatabase()
    r = gin.Default()
    routes.SetupRoutes(r)
}

func Handler(w http.ResponseWriter, req *http.Request) {
    r.ServeHTTP(w, req)
}
