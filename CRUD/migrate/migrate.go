package main

import (
	"github.com/gjustin40/CRUD/initalizers"
	"github.com/gjustin40/CRUD/models"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}
