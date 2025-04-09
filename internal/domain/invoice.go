package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	Id             string    `json:"id"`
	AccountId      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         Status    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreditCard struct {
	Number      string `json:"number"`
	Cvv         string `json:"cvv"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	HolderName  string `json:"holder_name"`
}

func NewInvoice(accountId string, amount float64, description string, paymentType string, card *CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		Id:             uuid.New().String(),
		AccountId:      accountId,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	var newStatus Status

	if randomSource.Float64() <= 0.2 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus

	return nil
}

func (i *Invoice) UpdateStatus(status Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = status
	i.UpdatedAt = time.Now()

	return nil
}
