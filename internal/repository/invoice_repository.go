package repository

import (
	"database/sql"

	"github.com/AlecDr/Full-Cycle-Payment-Gateway/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Save(invoice *domain.Invoice) error {
	query := `
	INSERT INTO invoices (id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(query, invoice.Id, invoice.AccountId, invoice.Amount, invoice.Status, invoice.Description, invoice.PaymentType, invoice.CardLastDigits, invoice.CreatedAt, invoice.UpdatedAt)
	return err
}

func (r *InvoiceRepository) FindById(id string) (*domain.Invoice, error) {
	query := `
	SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at
	FROM invoices
	WHERE id = $1
	`
	var invoice domain.Invoice

	err := r.db.QueryRow(query, id).Scan(
		&invoice.Id,
		&invoice.AccountId,
		&invoice.Amount,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentType,
		&invoice.CardLastDigits,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}

	return &invoice, nil
}

func (r *InvoiceRepository) FindByAccountId(accountId string) ([]*domain.Invoice, error) {
	query := `
	SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at
	FROM invoices
	WHERE account_id = $1
	`

	rows, err := r.db.Query(query, accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*domain.Invoice

	for rows.Next() {
		var invoice domain.Invoice
		err = rows.Scan(&invoice.Id, &invoice.AccountId, &invoice.Amount, &invoice.Status, &invoice.Description, &invoice.PaymentType, &invoice.CardLastDigits, &invoice.CreatedAt, &invoice.UpdatedAt)

		if err != nil {
			return nil, err
		}

		invoices = append(invoices, &invoice)
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}

	return invoices, nil
}

func (r *InvoiceRepository) UpdateStatus(invoice *domain.Invoice) error {
	query := `
	UPDATE invoices
	SET status = $1, updated_at = $2
	WHERE id = $3
	`
	rows, err := r.db.Exec(query, invoice.Status, invoice.UpdatedAt, invoice.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrInvoiceNotFound
	}

	return nil
}
