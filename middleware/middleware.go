package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mdzahid786/golang-graphql/helper"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		
		if authHeader == "" {
			next.ServeHTTP(w,r)
			return 
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(helper.JWTSecret), nil
		})
		
		if err != nil || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		// You can extract user ID or email from the claims
		userID := int32(claims["user_id"].(float64)) // or whatever your payload is
		ctx := context.WithValue(r.Context(), helper.UserCtxKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}