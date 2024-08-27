package user

import (
	"github.com/LUISEDOCCOR/api-mvc/database"
	"github.com/LUISEDOCCOR/api-mvc/models"
	"gorm.io/gorm"
)

func create(user models.User) *gorm.DB {
        result := database.DB.Create(user)
        return result
}
func getById(id uint) models.User {
        var user models.User
        database.DB.First(&user, id)
        return user
}
func getByEmail(email string) models.User {
        var user models.User
        database.DB.Where(&models.User{Email: email}).Find(&user)
        return user
}
