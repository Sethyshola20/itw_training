package repository

import (
	"billing_service/models"
	"context"
	"database/sql"
	"time"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, sub models.Subscription) error
	Get(ctx context.Context, id string) (models.Subscription, bool, error)
	List(ctx context.Context) ([]models.Subscription, error)
	Delete(ctx context.Context, id string) error
}

type subscriptionrepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) SubscriptionRepository {
	return &subscriptionrepository{db: db}
}

func (r *subscriptionrepository) Create(ctx context.Context, sub models.Subscription) error {
	query := `INSERT INTO subscriptions (id, user_id, plan, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		sub.ID, sub.UserID, sub.Plan, time.Now(), time.Now())
	return err
}

func (r *subscriptionrepository) Get(ctx context.Context, id string) (models.Subscription, bool, error) {
	query := `SELECT id, user_id, plan, created_at, updated_at FROM subscriptions WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var sub models.Subscription
	err := row.Scan(&sub.ID, &sub.UserID, &sub.Plan, &sub.CreatedAt, &sub.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.Subscription{}, false, nil
	}
	if err != nil {
		return models.Subscription{}, false, err
	}
	return sub, true, nil
}

func (r *subscriptionrepository) List(ctx context.Context) ([]models.Subscription, error) {
	query := `SELECT id, user_id, plan, created_at, updated_at FROM subscriptions`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		if err := rows.Scan(&sub.ID, &sub.UserID, &sub.Plan, &sub.CreatedAt, &sub.UpdatedAt); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

func (r *subscriptionrepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM subscriptions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
