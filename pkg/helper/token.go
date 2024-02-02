package helper

import (
	"context"
	"time"

	"github.com/FauzanAr/go-init/config"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/FauzanAr/go-init/pkg/wrapper"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

type AccessClaims struct {
	Claims
	jwt.StandardClaims
}

func GenerateAccessToken(ctx context.Context, claims Claims) (string, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)
	jwtClaims := AccessClaims{Claims: claims, StandardClaims: jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpired) * time.Hour).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := token.SignedString([]byte(cfg.Jwt.SecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateRefreshToken(ctx context.Context, claims Claims) (string, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)
	jwtClaims := AccessClaims{Claims: claims, StandardClaims: jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(cfg.Jwt.RefreshTokenExpired) * 24 * time.Hour).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := token.SignedString([]byte(cfg.Jwt.SecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func VerifyToken(ctx context.Context, tokenString string) (*AccessClaims, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)
	secretKey := []byte(cfg.Jwt.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, wrapper.UnauthorizedError("invalid token")
	}

	return claims, nil
}
