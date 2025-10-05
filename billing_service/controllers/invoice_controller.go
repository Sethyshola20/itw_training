package controllers

import (
	"billing_service/models"
	"billing_service/repository"
	"billing_service/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type InvoiceController struct {
	service *services.InvoiceService
}

func NewInvoiceController(repo repository.InvoiceRepository) *InvoiceController {
	return &InvoiceController{
		service: services.NewInvoiceService(repo),
	}
}

func (c *InvoiceController) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx := r.Context()
	id, err := c.service.CreateInvoiceService(ctx, invoice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	frontend := os.Getenv("FRONTEND_URL")
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Invoice créé avec succès, url:%s\n", frontend)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Invoice créé avec succès",
		"link":    fmt.Sprintf("%s/invoice/%v", frontend, id),
	})
}

func (c *InvoiceController) GetInvoice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	invoice, found, err := c.service.GetInvoiceService(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "Facture non trouvée",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "invoice récupéré avec succès",
		"invoice": invoice,
	})
}

func (c *InvoiceController) ListInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	invoices, found, err := c.service.ListInvoicesService(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "Aucune facture trouvée",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   http.StatusOK,
		"invoices": invoices,
	})
}

func (c *InvoiceController) GetInvoiceStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	ctx := r.Context()
	status, found, err := c.service.GetInvoiceStatusService(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "Aucune facture trouvée",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":         http.StatusOK,
		"invoice_status": status,
	})
}
