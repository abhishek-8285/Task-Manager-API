# Task Manager API (Go & Hexagonal Architecture)

Production-ready REST API built with Go, Gin, PostgreSQL, sqlc, and Docker.

---

## 🏗️ Architecture

This project follows **Hexagonal Architecture (Ports & Adapters)**:
- **`internal/domain`**: Pure business domain models (`User`).
- **`internal/ports`**: Interface contracts defining boundaries (`UserService`, `UserRepository`).
- **`internal/service`**: Core business logic layer (e.g. password hashing with `bcrypt`).
- **`internal/repository`**: DB adapter mapping type-safe `sqlc` queries to domain objects.
- **`internal/handler`**: HTTP adapter layer using Gin framework.

---

## 🚀 Quick Start

### Run Tests
```bash
make test
```

### Generate SQL Code with `sqlc`
```bash
make sqlc-generate
```

### Build & Run
```bash
make build
make run
```
