# Microservices Playground

This repository is a **playground for exploring microservice architecture** using multiple programming languages and technologies. The goal is to experiment with service separation, API gateways, containerization, databases, and inter-service communication in a hands-on, real-world scenario.

---

## 🧩 Project Overview

The project consists of several independent services:

| Service | Language / Framework | Purpose |
|---------|--------------------|---------|
| **Auth Service** | Python / Django | User authentication and management |
| **Billing Service** | Go | Manages subscriptions and invoices |
| **Payment Service** | Python | Handles payments (gRPC) |
| **API Gateway** | Go | Routes requests to the appropriate service, handles CORS and authentication |
| **UI** | Node.js / React | Frontend interface to interact with services |

Services communicate through HTTP REST APIs (except Payment via gRPC) and are designed to run independently in **Docker containers**, simulating a realistic microservice environment.

---

## ⚙️ Features

- Independent microservices in multiple programming languages
- API Gateway with CORS handling and route proxying
- PostgreSQL database for Billing service
- Dockerized environment for easy local development
- Example models for subscriptions and invoices
- Hot-reload support for Go services with Air
- Extensible structure for experimenting with new services

---

## 🛠️ Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.25+](https://golang.org/dl/) (for local builds if needed)
- [Node.js 18+](https://nodejs.org/) (for UI)

---

### Environment Variables

Create a `.env` file in the project root:

```env
# Billing Database
DB_HOST=billing-db
DB_PORT=5432
DB_USER=billing_user
DB_PASSWORD=supersecret
DB_NAME=billing

# API Gateway
AUTH_URL=http://auth:8000
BILLING_URL=http://billing:8090
UI_URL=http://localhost:3000
