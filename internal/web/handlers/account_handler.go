package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/domain"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/dto"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/service"
)

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.service.CreateAccount(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) FindByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if strings.TrimSpace(apiKey) == "" {
		http.Error(w, "X-API-Key header is required", http.StatusUnauthorized)
		return
	}

	account, err := h.service.FindByApiKey(apiKey)
	if err != nil && err != domain.ErrAccountNotFound {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err == domain.ErrAccountNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if strings.TrimSpace(id) == "" {
		http.Error(w, "id query parameter is required", http.StatusBadRequest)
		return
	}

	account, err := h.service.FindById(id)
	if err != nil && err != domain.ErrAccountNotFound {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err == domain.ErrAccountNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
