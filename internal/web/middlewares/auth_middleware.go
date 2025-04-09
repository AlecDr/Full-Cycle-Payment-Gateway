package middlewares

import (
	"net/http"
	"strings"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/domain"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/service"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{accountService: accountService}
}

func (m *AuthMiddleware) Run(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if strings.TrimSpace(apiKey) == "" {
			http.Error(w, "X-API-Key header not found", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByApiKey(apiKey)
		if err != nil {
			switch err {
			case domain.ErrAccountNotFound:
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		next.ServeHTTP(w, r)

	})
}
