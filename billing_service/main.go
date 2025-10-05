package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
	"billing_service/controllers"
	"billing_service/middleware"
	"billing_service/repository"
	"billing_service/routes"
)

func main() {
	mux := http.NewServeMux()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	subRepo := repository.NewSubscriptionRepository(db)
	invRepo := repository.NewInvoiceRepository(db)

	subCtrl := controllers.NewSubscriptionController(subRepo)
	invCtrl := controllers.NewInvoiceController(invRepo)

	routes.SetupSubscriptionRoutes(mux, subCtrl)
	routes.SetupInvoiceRoutes(mux, invCtrl)

	var handler http.Handler = mux
	handler = middleware.Logger(handler)
	handler = middleware.CORS(handler)

	log.Println("Billing Service listening on :8002")
	log.Fatal(http.ListenAndServe(":8002", handler))
}
