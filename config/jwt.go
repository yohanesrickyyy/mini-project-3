package config

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	UserID int
	jwt.RegisteredClaims
}

func GetJWTKey() []byte {
	return []byte(GetEnv("JWT_KEY"))
}
