package service

import (
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/domain"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/dto"
	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/repository"
)

type InvoiceService struct {
	repository     *repository.InvoiceRepository
	accountService *AccountService
}

func NewInvoiceService(repository *repository.InvoiceRepository, accountService *AccountService) *InvoiceService {
	return &InvoiceService{
		repository:     repository,
		accountService: accountService,
	}
}

func (s *InvoiceService) CreateInvoice(input *dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	account, err := s.accountService.FindByApiKey(input.ApiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(account.Id, input)
	if err != nil {
		return nil, err
	}

	err = invoice.Process()
	if err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.ApiKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	err = s.repository.Save(invoice)
	if err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) FindById(id string, apiKey string) (*dto.InvoiceOutput, error) {
	account, err := s.accountService.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	if invoice.AccountId != account.Id {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) ListByAccountId(accountId string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.repository.FindByAccountId(accountId)
	if err != nil {
		return nil, err
	}

	invoicesOutput := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		invoicesOutput[i] = dto.FromInvoice(invoice)
	}

	return invoicesOutput, nil
}

func (s *InvoiceService) ListByAccountApiKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	account, err := s.accountService.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccountId(account.Id)
}
