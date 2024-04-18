package initializer

import "github.com/goropencho/golang-gin-auth/models"

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
}
