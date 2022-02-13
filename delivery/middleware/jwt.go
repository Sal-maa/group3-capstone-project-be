package middleware

import (
	_config "capstone/be/config"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = _config.GetConfig()

func JWTMiddleWare() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.JWT_secret),
	})
}

func CreateToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_secret))
}

func ValidateToken(c echo.Context) bool {
	login := c.Get("user").(*jwt.Token)

	return login.Valid
}

func ExtractId(c echo.Context) int {
	login := c.Get("user").(*jwt.Token)

	claims := login.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))

	return id
}
