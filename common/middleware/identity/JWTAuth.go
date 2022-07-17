package identity

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth struct {
	signingKey string
}

func NewAuth(key string) *Auth {
	return &Auth{
		signingKey: key,
	}
}

func (a *Auth) CreateToken(userID string, tokenTTL uint64) (string, string, int64, error) {
	mySigningKey := []byte(a.signingKey)

	// generate access_token
	duration := time.Duration(tokenTTL) * time.Second
	expirationTime := time.Now().Add(duration)

	claims := jwt.MapClaims{
		"id":  uuid.New().String(),
		"sub": userID,
		"exp": expirationTime.Unix(),
		"iss": "common-service",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenStr, err := accessToken.SignedString(mySigningKey)
	if err != nil {
		return "", "", 0, err
	}

	// generate refresh_token
	claims["id"] = uuid.New().String()
	claims["exp"] = time.Now().Add((time.Duration(tokenTTL*2) * time.Second))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenStr, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return "", "", 0, err
	}

	return accessTokenStr, refreshTokenStr, int64(duration.Seconds()), err
}

func (a *Auth) ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.New(codes.InvalidArgument, fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])).Err()
		}

		return []byte(a.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, status.New(codes.Unauthenticated, "unauthorized").Err()
	}

	return claims, nil
}
