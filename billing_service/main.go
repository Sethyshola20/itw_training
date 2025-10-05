package main

import (
	"database/sql"
	"fmt"
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

	// Read DB connection info from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("All DB_* environment variables are required")
	}

	// Construct PostgreSQL connection string
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open DB connection:", err)
	}
	defer db.Close()

	// Ping DB to ensure connectivity
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	// Initialize repositories
	subRepo := repository.NewSubscriptionRepository(db)
	invRepo := repository.NewInvoiceRepository(db)

	// Initialize controllers
	subCtrl := controllers.NewSubscriptionController(subRepo)
	invCtrl := controllers.NewInvoiceController(invRepo)

	// Setup routes
	routes.SetupSubscriptionRoutes(mux, subCtrl)
	routes.SetupInvoiceRoutes(mux, invCtrl)

	// Apply middleware
	var handler http.Handler = mux
	handler = middleware.Logger(handler)
	handler = middleware.CORS(handler)

	log.Println("Billing Service listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", handler))
}
