package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//JWTSecret := os.Getenv("JWT_SECRET")
		JWTSecret := "G3p3N9OsoQ8tOog5V6jLCdSrZPxmuj7AW5FSvJHKWDC0K_kGr1Tl4aL80aRXpHWR2H3WE0fdYquE5rSakHRhMsu9QAgbPg=="
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		jwtToken := authHeader[1]

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWTSecret), nil
		})

		if token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			fmt.Print(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email, _ := claims["email"].(string)
			ID, _ := claims["ID"].(string)

			ctx := context.WithValue(r.Context(), "email", email)
			ctx = context.WithValue(ctx, "ID", ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Print(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}

	})
}
