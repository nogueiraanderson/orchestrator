package http

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// jwtAudience is the audience a bearer token must name to be granted write access.
const jwtAudience = "orchestrator"

// jwtAudienceOK reports whether the bearer token in req names jwtAudience.
func jwtAudienceOK(req *http.Request) bool {
	raw := strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")
	claims := jwt.MapClaims{}
	if _, _, err := new(jwt.Parser).ParseUnverified(raw, claims); err != nil {
		return false
	}
	return claims.VerifyAudience(jwtAudience, true)
}
