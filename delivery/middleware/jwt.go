package middleware

import (
	_config "capstone/be/config"
	_common "capstone/be/delivery/common"
	"net/http"

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
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			switch e {
			case middleware.ErrJWTMissing:
				return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "missing or malformed jwt"))
			default:
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid or expired jwt"))
			}
		},
	})
}

func CreateToken(id int, role string) (string, int64, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	expire := time.Now().Add(time.Hour * 1).Unix()
	claims["exp"] = expire

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := _token.SignedString([]byte(config.JWT_secret))

	return token, expire, err
}

func ExtractId(c echo.Context) int {
	login := c.Get("user").(*jwt.Token)

	claims := login.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))

	return id
}

func ExtractRole(c echo.Context) string {
	login := c.Get("user").(*jwt.Token)

	claims := login.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	return role
}
