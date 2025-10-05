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

type SubscriptionController struct {
	service *services.SubscriptionService
}

func NewSubscriptionController(repo repository.SubscriptionRepository) *SubscriptionController {
	return &SubscriptionController{
		service: services.NewSubscriptionService(repo),
	}
}

func (c *SubscriptionController) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var subscription models.Subscription
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx := r.Context()
	id, err := c.service.CreateSubscriptionService(ctx, subscription)
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
	fmt.Printf("Subscription créé avec succès, url:%s\n", frontend)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Subscription créé avec succès",
		"link":    fmt.Sprintf("%s/subscription/%v", frontend, id),
	})
}

func (c *SubscriptionController) GetSubscription(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	subscription, found, err := c.service.GetSubscriptionService(ctx, id)
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
			"message": "Subscription non trouvé",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":       http.StatusOK,
		"message":      "Subscription récupéré avec succès",
		"subscription": subscription,
	})
}

func (c *SubscriptionController) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	subscriptions, found, err := c.service.ListSubscriptionsService(ctx)
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
			"message": "Aucun abonnement trouvé",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        http.StatusOK,
		"subscriptions": subscriptions,
	})
}

func (c *SubscriptionController) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	err := c.service.DeleteSubscriptionService(ctx, id) 
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "Subscription non trouvé",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Subscription supprimé avec succès",
	})
}
