package types

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Auth struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	UserId uuid.UUID `json:"uid"`
	jwt.RegisteredClaims
}
