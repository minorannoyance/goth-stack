package middlewares

import (
	"crypto/subtle"
	"log"
	"os"

	// https://echo.labstack.com/docs/middleware/jwt
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	})

}

func JWTAuth() echo.MiddlewareFunc {
	// read secret jwtSecretKey from file
	jwtSecretKey, err := os.ReadFile("secret.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Printf("Key: %s", jwtSecretKey)
	}
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecretKey),
	})
}

func CSRF() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	})
}
