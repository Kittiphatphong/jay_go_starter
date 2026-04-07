package controllers

import (
	"go_starter/config"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

type AuthWeb struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func GenerateTokenWeb(user AuthWeb) (string, error) {
	// Extract the credentials from the request body
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"id":       strconv.Itoa(int(user.ID)),
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Env("auth.secretWeb")))
	if err != nil {
		return "", err
	}
	return t, nil
}

type AuthApi struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func GenerateTokenApi(partner AuthApi) (string, error) {
	// Extract the credentials from the request body
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"id":       strconv.Itoa(int(partner.ID)),
		"username": partner.Username,
		"name":     partner.Name,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Env("auth.secretApi")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["username"].(string)
	return c.SendString("Welcome 👋" + email)
}
