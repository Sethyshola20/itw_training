package repository

import (
	"billing_service/models"
	"context"
	"database/sql"
	"time"
)

type Invoice struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InvoiceRepository interface {
	Create(ctx context.Context, sub models.Invoice) error
	Get(ctx context.Context, id string) (models.Invoice, bool, error)
	List(ctx context.Context) ([]models.Invoice, error)
	Delete(ctx context.Context, id string) error
}

type Invoicerepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) InvoiceRepository {
	return &Invoicerepository{db: db}
}

func (r *Invoicerepository) Create(ctx context.Context, sub models.Invoice) error {
	query := `INSERT INTO subscriptions (id, user_id, plan, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		sub.ID, sub.Amount, sub.Status, time.Now(), time.Now())
	return err
}

func (r *Invoicerepository) Get(ctx context.Context, id string) (models.Invoice, bool, error) {
	query := `SELECT id, user_id, plan, created_at, updated_at FROM subscriptions WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var sub models.Invoice
	err := row.Scan(&sub.ID, &sub.Amount, &sub.Status, &sub.CreatedAt, &sub.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.Invoice{}, false, nil
	}
	if err != nil {
		return models.Invoice{}, false, err
	}
	return sub, true, nil
}

func (r *Invoicerepository) List(ctx context.Context) ([]models.Invoice, error) {
	query := `SELECT id, user_id, plan, created_at, updated_at FROM subscriptions`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Invoice
	for rows.Next() {
		var sub models.Invoice
		if err := rows.Scan(&sub.ID, &sub.Amount, &sub.Status, &sub.CreatedAt, &sub.UpdatedAt); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

func (r *Invoicerepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM subscriptions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
