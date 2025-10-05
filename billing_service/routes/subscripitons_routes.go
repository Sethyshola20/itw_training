package routes

import (
	"net/http"

	"billing_service/controllers"
)

func SetupSubscriptionRoutes(mux *http.ServeMux, billingController *controllers.SubscriptionController) {
	mux.HandleFunc("GET /api/billing/subscriptions", billingController.ListSubscriptions)
	mux.HandleFunc("POST /api/billing/subscriptions", billingController.CreateSubscription)
}
