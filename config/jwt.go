package config

import "github.com/golang-jwt/jwt/v4"

var SecretKey = []byte("rahasia-super-aman")

type JWTClaim struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
} 