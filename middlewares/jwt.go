package middlewares

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"../structs"
	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey    = []byte(os.Getenv("JTW_SECRET_KEY"))
	configExpired = strings.Split(os.Getenv("JWT_TOKEN_EXPIRED"), ",")
)

//Claims jwt claims  structure
type Claims struct {
	ID uint
	jwt.StandardClaims
}

//JWTMiddleware jwt middleware
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	var (
		responseStruct structs.ResponseStruct
	)
	return func(w http.ResponseWriter, r *http.Request) {
		result := &responseStruct
		w.Header().Set("Content-Type", "application/json")
		jwtTokenString := r.Header.Get("Authorization")
		if len(jwtTokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			result.Status = false
			result.Message = "Authentication failure"
			result.Result = nil
			json.NewEncoder(w).Encode(result)
			return
		}
		jwtTokenString = strings.Replace(jwtTokenString, "Bearer ", "", 1)
		claims := &Claims{}
		if tkn, err := jwt.ParseWithClaims(jwtTokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		}); err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			result.Status = false
			result.Message = "Authentication failure"
			result.Result = nil
			json.NewEncoder(w).Encode(result)
			return
		}
		next(w, r)
	}
}

//GenerateJWT generate token key
func GenerateJWT(data structs.Users) (string, error) {
	expired := make([]int, len(configExpired))
	for i := range expired {
		configExpired[i] = strings.ReplaceAll(configExpired[i], " ", "")
		expired[i], _ = strconv.Atoi(configExpired[i])
	}

	claims := &Claims{
		ID: data.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(expired[0], expired[1], expired[2]).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
