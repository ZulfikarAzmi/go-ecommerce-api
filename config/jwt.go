package config

import "github.com/golang-jwt/jwt/v4"

var SecretKey = []byte("rahasia-super-aman")

// JWTClaim adalah struktur untuk klaim JWT
type JWTClaim struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    jwt.RegisteredClaims
} 