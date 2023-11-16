package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("ef74HBH3nf34x34ry7HBSAsdaaDMXdasdUHNghn327zr2")

type JWTClaims struct {
	jwt.RegisteredClaims
	Username	string	`json:"username"`
	Email		string	`json:"email"`
}