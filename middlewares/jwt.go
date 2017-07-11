package middlewares

import "github.com/dgrijalva/jwt-go"

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	jwt.StandardClaims
}
