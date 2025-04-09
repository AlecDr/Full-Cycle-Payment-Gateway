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
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	PaymentType     string  `json:"payment_type"`
	CardNumber      string  `json:"card_number"`
	Cvv             string  `json:"cvv"`
	ExpirationMonth int     `json:"expiration_month"`
	ExpirationYear  int     `json:"expiration_year"`
	HolderName      string  `json:"holder_name"`
}

type InvoiceOutput struct {
	Id             string        `json:"id"`
	AccountId      string        `json:"account_id"`
	Amount         float64       `json:"amount"`
	Status         domain.Status `json:"status"`
	Description    string        `json:"description"`
	PaymentType    string        `json:"payment_type"`
	CardLastDigits string        `json:"card_last_digits"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
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
