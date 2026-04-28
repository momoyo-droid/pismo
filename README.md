# Pismo API

REST API developed in Go for account and financial transaction management.

## Technologies Used

- Go
- Gin
- GORM
- PostgreSQL
- Docker
- Swagger
- Unit Tests
- GitHub Actions

---

# Prerequisites

Before running the project, make sure you have the following installed:

- Go
- Docker
- Docker Compose
- Make

---

# Running the Application

## 1. Clone the repository

```bash
git clone https://github.com/momoyo-droid/pismo.git
cd pismo
```

---

## 2. Run the application

```bash
make build
```

This command will:

- stop old containers
- build the application
- start the API and PostgreSQL containers

If one of the containers dont start, restart using `docker compose restart <service name>`.
Service name can be `api` or `postgres`

---

# Swagger / API Documentation

After starting the application, the Swagger documentation will be available at:

```text
http://localhost:3000/swagger/index.html
```

To update the Swagger documentation:

```bash
make swagger
```

---

# Running Tests

## Run unit tests

```bash
make test
```

---

# Code Quality Tools

## Format code

```bash
make fmt
```

---

## Static analysis

```bash
make vet
```

---

## Run linter

```bash
make lint
```

---

## Vulnerability check

```bash
make vuln
```

---

## Run full audit

Runs:

- formatting
- vet
- lint
- vulnerability analysis
- tests

```bash
make audit
```

---

# Project Structure

```text
api/
├── cmd/
    ├── app/
├── internal/
|   ├── config/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   ├── service/
│   └── utils/
├── docs/
```

---

# Features

## Accounts

- Create account
- Get account by ID

## Transactions

- Create financial transactions

---

# Author

Ana Oliveira
