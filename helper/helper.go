package helper

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWTSecret = []byte("your-secret-key")
type contextKey string
const UserCtxKey = contextKey("user_id")

func GenerateJWT(userID int32) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func GetUserIDFromContext(ctx context.Context)(int32, error) {
	userID := ctx.Value(UserCtxKey)
	if userID == nil {
		return 0, errors.New("Unauthentiated")
	}
	return userID.(int32), nil
}