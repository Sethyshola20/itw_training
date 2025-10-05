package routes

import (
	"net/http"

	"billing_service/controllers"
)

func SetupInvoiceRoutes(mux *http.ServeMux, billingController *controllers.InvoiceController) {
	mux.HandleFunc("GET /api/billing/invoices", billingController.ListInvoices)
	mux.HandleFunc("GET /api/billing/invoices/{id}", billingController.GetInvoice)
	mux.HandleFunc("GET /api/billing/invoices/{id}/status", billingController.GetInvoiceStatus)
}
