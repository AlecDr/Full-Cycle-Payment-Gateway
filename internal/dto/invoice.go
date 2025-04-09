package dto

import (
	"time"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	ApiKey          string
	Amount          float64
	Description     string
	PaymentType     string
	CardNumber      string
	Cvv             string
	ExpirationMonth int
	ExpirationYear  int
	HolderName      string
}

type InvoiceOutput struct {
	Id             string
	AccountId      string
	Amount         float64
	Status         domain.Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func ToInvoice(accountId string, input *CreateInvoiceInput) (*domain.Invoice, error) {
	creditCard := domain.CreditCard{
		Number:      input.CardNumber,
		Cvv:         input.Cvv,
		ExpiryMonth: input.ExpirationMonth,
		ExpiryYear:  input.ExpirationYear,
		HolderName:  input.HolderName,
	}
	return domain.NewInvoice(accountId, input.Amount, input.Description, input.PaymentType, &creditCard)
}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		Id:             invoice.Id,
		AccountId:      invoice.AccountId,
		Amount:         invoice.Amount,
		Status:         invoice.Status,
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
