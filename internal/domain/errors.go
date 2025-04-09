package domain

import "errors"

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrDuplicatedApiKey    = errors.New("duplicated api key")
	ErrInvoiceNotFound     = errors.New("invoice not found")
	ErrUnauthorizedAccess  = errors.New("unauthorized access")
	ErrInsufficientBalance = errors.New("insufficient balance")

	ErrInvalidAmount = errors.New("amount must be greater than 0")
	ErrInvalidStatus = errors.New("invalid status")
)
