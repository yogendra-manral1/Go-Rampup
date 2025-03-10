package middlewares

import (
	"Go-Rampup/config"
	"errors"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func PanicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.AbortWithError(http.StatusInternalServerError, nil)
			}
		}()
		c.Next()
	}
}

func Authorization(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization")
	if tokenString == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	token, err := VerifyToken(tokenString)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	email := token.Claims.(jwt.MapClaims)["email"].(string)
	context.Set("email", email)
}

func CreateToken(context *gin.Context) {
	if _, exists := context.Get("email"); !exists {
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": context.MustGet("email"),
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JWTSecretKey))
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JWTSecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
