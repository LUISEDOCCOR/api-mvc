package utils

import (
	"os"

	"github.com/LUISEDOCCOR/api-mvc/models"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetJwtPassword() string {
	return os.Getenv("jwt_password")
}

func CreateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//e == error and t == token as string
	t, e := token.SignedString([]byte(GetJwtPassword()))

	return t, e
}

func CreateResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"msg": message,
	})
}

func HashPassword(user *models.User) {
	hash, _ := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	user.Password = hash
}
