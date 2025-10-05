package services

import (
	"billing_service/models"
	"billing_service/repository"
	"context"

	"github.com/google/uuid"
)

type InvoiceService struct {
	repo repository.InvoiceRepository 
}

func NewInvoiceService(repo repository.InvoiceRepository) *InvoiceService {
	return &InvoiceService{repo: repo}
}

func (s *InvoiceService) CreateInvoiceService(ctx context.Context, invoice models.Invoice) (string, error) {
	id := uuid.New().String()
	invoice.ID = id 

	err := s.repo.Create(ctx, invoice)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *InvoiceService) GetInvoiceService(ctx context.Context, id string) (models.Invoice, bool, error) {
	invoice, found, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.Invoice{}, false, err
	}
	if !found {
		return models.Invoice{}, false, nil
	}
	return invoice, true, nil
}

func (s *InvoiceService) GetInvoiceStatusService(ctx context.Context, id string) (string, bool, error) {
	invoice, found, err := s.repo.Get(ctx, id)
	if err != nil {
		return "", false, err
	}
	if !found {
		return "", false, nil
	}
	status := invoice.Status
	return status, true, nil
}

func (s *InvoiceService) ListInvoicesService(ctx context.Context) ([]models.Invoice, bool, error) {
	invoices, err := s.repo.List(ctx)
	if err != nil {
		return nil, false, err
	}
	if len(invoices) == 0 {
		return nil, false, nil
	}
	return invoices, true, nil
}
