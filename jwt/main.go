package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gjustin40/jwt/controllers"
	"github.com/gjustin40/jwt/initalizers"
	"github.com/gjustin40/jwt/middleware"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnetToDb()
	initalizers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
