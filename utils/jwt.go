package utils

import (
	"os"
	"time"

	"github.com/cla-github/file_manager/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJwt(uid uuid.UUID) (string, error) {
	claims := types.Claims{
		UserId: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(jwtSecret)
}
