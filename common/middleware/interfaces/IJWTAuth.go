package interfaces

import "github.com/golang-jwt/jwt"

// IJWTAuth : interface
type IJWTAuth interface {
	ValidateToken(token string) (jwt.MapClaims, error)
}
