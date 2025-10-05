package services

import (
	"billing_service/models"
	"billing_service/repository"
	"context"

	"github.com/google/uuid"
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscriptionService(ctx context.Context, subscription models.Subscription) (string, error) {
	id := uuid.New().String()
	subscription.ID = id 

	// Save in repo
	err := s.repo.Create(ctx, subscription)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *SubscriptionService) GetSubscriptionService(ctx context.Context, id string) (models.Subscription, bool, error) {
	subscription, found, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.Subscription{}, false, err
	}
	if !found {
		return models.Subscription{}, false, nil
	}
	return subscription, true, nil
}

func (s *SubscriptionService) ListSubscriptionsService(ctx context.Context) ([]models.Subscription, bool, error) {
	subscriptions, err := s.repo.List(ctx)
	if err != nil {
		return nil, false, err
	}
	if len(subscriptions) == 0 {
		return nil, false, nil
	}
	return subscriptions, true, nil
}

func (s *SubscriptionService) DeleteSubscriptionService(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
