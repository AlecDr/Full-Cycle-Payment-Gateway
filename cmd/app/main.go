package main

import (
	"log"
	"os"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/cmd/app/db"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/repository"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/service"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/web/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := db.NewDb()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	invoiceRepository := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepository, accountService)

	server := server.NewServer(os.Getenv("HTTP_PORT"), accountService, invoiceService)

	err := server.Start()
	if err != nil {
		log.Fatal("Error starting server")
	}
}
