package server

import (
	"fmt"
	"net/http"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/service"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(port string, accountService *service.AccountService) *Server {
	router := chi.NewRouter()

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	server := &Server{
		router:         router,
		server:         httpServer,
		accountService: accountService,
		port:           port,
	}

	server.SetupRoutes()

	return server
}

func (s *Server) SetupRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Route("/accounts", func(r chi.Router) {
		r.Post("/", accountHandler.CreateAccount)
		r.Get("/", accountHandler.FindByApiKey)
		r.Get("/{id}", accountHandler.FindById)
	})
}

func (s *Server) Start() error {
	fmt.Println("Starting server on port", s.port)
	return s.server.ListenAndServe()
}
