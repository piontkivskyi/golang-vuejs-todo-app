package controllers

import (
	"net/http"
	"time"
	"todo-app/middlewares"
	"todo-app/models"
	"todo-app/models/connection"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetToken return jwt by given username and password
func GetToken(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := new(models.User)
	// try to get task by given id
	if res := connection.DB.Where(&models.User{Username: username}).First(&user); res.Error != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
		// Set custom claims
		claims := &middlewares.JwtCustomClaims{
			Name:   user.Name,
			UserID: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("get from env in future"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	return c.NoContent(http.StatusUnauthorized)
}
