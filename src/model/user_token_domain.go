package model

import (
	"os"
	"strings"
	"time"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.email,
		"age":   ud.age,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 240).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError("Error generating token")
	}

	return tokenString, nil
}

func RemoveBearerFromToken(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}

	return token
}

func VerifyToken(token string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token = RemoveBearerFromToken(token)

	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)

	if !ok || !tkn.Valid {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	return &userDomain{
		ID:    claims["id"].(string),
		email: claims["email"].(string),
		age:   int8(claims["age"].(float64)),
		name:  claims["name"].(string),
	}, nil
}

func AuthorizationMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token := c.Request.Header.Get("Authorization")

	token = RemoveBearerFromToken(token)

	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)

	if !ok || !tkn.Valid {
		errRest := rest_err.NewUnauthorizedError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := &userDomain{
		ID:    claims["id"].(string),
		email: claims["email"].(string),
		age:   int8(claims["age"].(float64)),
		name:  claims["name"].(string),
	}

	logger.Info("Middleware catch: ", zap.String("user_id", userDomain.GetID()))
}
