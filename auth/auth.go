package auth

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type AuthMiddleware struct {
	authToken string
}

func NewAuthMiddleware(token string) *AuthMiddleware {
	return &AuthMiddleware{
		authToken: token,
	}
}

func (a *AuthMiddleware) Handler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if h, ok := r.Header["X-Auth-Token"]; ok {
		if h[0] == a.authToken {
			next(rw, r)
			return
		}
	}

	log.Warnf("unauthorized request: addr=%s", r.RemoteAddr)

	rw.WriteHeader(http.StatusUnauthorized)
	rw.Write([]byte("unauthorized\n"))
}
