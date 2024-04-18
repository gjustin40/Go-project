package initalizers

import "github.com/gjustin40/jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
