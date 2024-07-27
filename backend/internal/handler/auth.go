package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/m-Chetan/go-shawty/database"
	"github.com/m-Chetan/go-shawty/internal/helpers"
	"github.com/m-Chetan/go-shawty/internal/model"
)

var jwtKey = []byte("Secret123")

func Login(c *fiber.Ctx) error {
	db := database.DB

	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Request"})
	}

	var existingUser model.User

	res := db.Where("email=?", user.Email).First(&existingUser)

	if res.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User does not exists"})
	}

	passwordMatch := helpers.CompareHashPassword(existingUser.Password, user.Password)

	if !passwordMatch {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Password"})
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: existingUser.Email, ExpiresAt: expirationTime.Unix()})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = expirationTime

	c.Cookie(cookie)
	return c.Status(200).JSON(fiber.Map{"success": "User logged in successfully"})

}

func Signup(c *fiber.Ctx) error {
	db := database.DB

	var user model.User
	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var existingUser model.User

	res := db.Where("email=?", existingUser.Email).First(&existingUser)

	if res.Error == nil {
		return c.Status(404).JSON(fiber.Map{"error": "User already exists"})
	}

	hashedPassword, err := helpers.GenerateHashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate hash of password"})
	}

	user.Password = hashedPassword

	res = db.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return c.Status(200).JSON(fiber.Map{"success": "User created"})
}
