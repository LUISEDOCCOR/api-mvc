package user_model

import (
	"github.com/LUISEDOCCOR/api-mvc/database"
	"github.com/LUISEDOCCOR/api-mvc/models"
)

func Create(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}
func GetById(id uint) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return user, result.Error
}
func GetByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where(&models.User{Email: email}).Find(&user)
	return user, result.Error
}
