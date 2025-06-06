package utils

import (
	"net/http"
	"strings"
	"user-management/models"
	"user-management/repository"
)

func ChainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)	
	}

	return h
}

func AuthenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if authHeader == "" || token == authHeader {
			NoToken(w)
			return
		}

		userID, err := VerifyToken(token)
		if err != nil {
			LogError(r, err)
			InvalidToken(w)
			return
		}

		if _, err = repository.GetUserByID(userID); err != nil {
			LogError(r, err)
			InternalServerError(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if authHeader == "" || token == authHeader {
			NoToken(w)
			return
		}

		userID, err := VerifyToken(token)
		if err != nil {
			LogError(r, err)
			InvalidToken(w)
			return
		}

		user, err := repository.GetUserByID(userID)
		if err != nil {
			LogError(r, err)
			InternalServerError(w)
			return
		}

		if user.Role != models.RoleAdmin {
			Forbidden(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}