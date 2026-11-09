package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtNotExpired reports whether the bearer token in req carries an exp claim in the future.
func jwtNotExpired(req *http.Request) bool {
	raw := strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")
	claims := jwt.MapClaims{}
	if _, _, err := jwt.NewParser().ParseUnverified(raw, claims); err != nil {
		return false
	}
	exp, err := claims.GetExpirationTime()
	if err != nil || exp == nil {
		return false
	}
	return exp.After(time.Now())
}
