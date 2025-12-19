# Auth Service (Go)

A backend ewallet service built with **Go**, **Gin**, and **PostgreSQL**, designed using a clean and scalable architecture.  
This project is intended as a learning and portfolio project, focusing on backend fundamentals and best practices.

---

## ✨ Features (Current)
- Application bootstrap with structured logging (Zap)
- PostgreSQL connection with health check
- Clean project structure (repository, service/usecase, handler, route)
- Database migrations (local development)
- Database seeding for local environment
- Centralized route setup with Gin

---

## 🛠️ Tech Stack
- **Language**: Go
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM / Driver**: pgx
- **Logging**: Zap
- **Architecture**: Clean Architecture (Repository → Service/Usecase → Handler)
- **Environment Config**: `.env`

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── dto/
│   │   ├── handler/
│   │   ├── model/
│   │   ├── repository/
│   │   └── route/
│   │   └── service/
│   ├── apperror/
│   ├── config/
│   │   ├── contract.go
│   │   ├── migration.go
│   │   └── zap_sugared_logger.go
│   ├── dto/
│   ├── middleware/
│   └── util/
├── migrations/
│   ├── 000_drop_tables.sql
│   ├── 001_schema.sql
│   └── 002_seeder.sql
├── pkg/
│   └── connect_db.go
├── .env.example
├── go.mod
└── README.md
