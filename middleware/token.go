package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte(os.Getenv("JWT_token"))

// Token Generate function
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Token verify function
func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	userID := fmt.Sprintf("%v", claims["user_id"])
	return userID, nil
}

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			}

			tokenString := parts[1]
			userID, err := VerifyToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}
			fmt.Println("userID----", userID)
			// Save userID in context
			c.Set("userID", userID)

			return next(c)
		}
	}
}
