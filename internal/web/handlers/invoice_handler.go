package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/domain"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/dto"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/service"
	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	service *service.InvoiceService
}

func NewInvoiceHandler(service *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if strings.TrimSpace(apiKey) == "" {
		http.Error(w, "X-API-Key header is required", http.StatusUnauthorized)
		return
	}

	input.ApiKey = apiKey

	invoice, err := h.service.CreateInvoice(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if strings.TrimSpace(id) == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if strings.TrimSpace(apiKey) == "" {
		http.Error(w, "X-API-Key header is required", http.StatusUnauthorized)
		return
	}

	invoice, err := h.service.FindById(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if strings.TrimSpace(apiKey) == "" {
		http.Error(w, "X-API-Key header is required", http.StatusUnauthorized)
		return
	}

	invoices, err := h.service.ListByAccountApiKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoices)
}
