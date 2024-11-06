package authmiddleware

import (
	"context"
	"encoding/json"
	"goph-keeper/internal/modules/auth/authservices/authjwtservice"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// Тип для ключей контекста.
type contextKey string

const Claims contextKey = "claims"

func Authentication(next http.Handler, secretKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jwtBearer := r.Header.Get("Authorization")

		strArr := strings.Split(jwtBearer, " ")
		if len(strArr) != 2 { //nolint:gomnd // 2 части - Bearer + JWT
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Authorization header format must be 'Bearer <token>'",
			})
			return
		}

		token, err := authjwtservice.ParseAndValidateToken(strArr[1], secretKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Token validation failed",
			})
			return
		}

		ctx := context.WithValue(r.Context(), Claims, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
