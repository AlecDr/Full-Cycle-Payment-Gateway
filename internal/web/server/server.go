package server

import (
	"fmt"
	"net/http"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/service"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/web/handlers"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(port string, accountService *service.AccountService, invoiceService *service.InvoiceService) *Server {
	router := chi.NewRouter()

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	server := &Server{
		router:         router,
		server:         httpServer,
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}

	server.SetupRoutes()

	return server
}

func (s *Server) SetupRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middlewares.NewAuthMiddleware(s.accountService)

	s.router.Route("/accounts", func(r chi.Router) {
		r.Post("/", accountHandler.CreateAccount)
		r.Get("/", accountHandler.FindByApiKey)
		r.Get("/{id}", accountHandler.FindById)
	})

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Run)

		r.Route("/invoices", func(r chi.Router) {
			r.Post("/", invoiceHandler.CreateInvoice)
			r.Get("/", invoiceHandler.ListByAccount)
			r.Get("/{id}", invoiceHandler.FindById)
		})
	})
}

func (s *Server) Start() error {
	fmt.Println("Starting server on port", s.port)
	return s.server.ListenAndServe()
}
