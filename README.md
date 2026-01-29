# TODO-MIS

A RESTful API service for TODO item management built with Go, following Hexagonal Architecture (Ports and Adapters) principles.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Running with Docker Compose](#running-with-docker-compose)
  - [Running Locally](#running-locally)
- [Testing](#testing)
- [Building](#building)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Development](#development)

## Features

- ✅ CRUD operations for TODO items
- ✅ OAuth-based third-party authentication (Google, GitHub, Facebook)
- ✅ JWT-based authorization
- ✅ Hexagonal Architecture for clean separation of concerns
- ✅ Dependency injection with Uber FX
- ✅ MySQL database with GORM
- ✅ Docker and Docker Compose support
- ✅ Health check endpoint
- ✅ Structured logging with Zap

## Architecture

This project follows **Hexagonal Architecture** (also known as Ports and Adapters):

```
├── adapter/          # Adapters (Infrastructure Layer)
│   ├── driven/       # Driven adapters (outbound)
│   │   ├── auth/     # OAuth providers
│   │   └── persistence/  # Database implementations
│   └── driving/      # Driving adapters (inbound)
│       └── api/      # HTTP handlers
├── application/      # Application services (Use Cases)
├── domain/          # Domain layer (Business Logic)
│   ├── entity/      # Domain entities
│   ├── todo/        # TODO domain logic
│   └── auth/        # Auth domain logic
├── server/          # Server infrastructure
└── common/          # Shared utilities
```

## Tech Stack

- **Language**: Go 1.25.5
- **Web Framework**: Gin
- **Dependency Injection**: Uber FX
- **ORM**: GORM
- **Database**: MySQL 8.0
- **Authentication**: JWT (golang-jwt/jwt)
- **Logging**: Zap
- **Testing**: Testify, GoMock
- **Containerization**: Docker, Docker Compose

## Prerequisites

### For Docker Compose (Recommended)
- Docker 20.10+
- Docker Compose 2.0+

### For Local Development
- Go 1.25.5+
- MySQL 8.0+
- Make (optional)

## Getting Started

### Running with Docker Compose

This is the recommended way to run the application with all dependencies.

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd TODO-MIS
   ```

2. **Configure environment variables**

   Create a `.env` file in the project root:
   ```env
   # Application
   APP_ENV=dev
   PORT=8080
   
   # MySQL (for local development)
   MYSQL_DSN=root:root@tcp(localhost:3306)/todo-mis?charset=utf8mb4&parseTime=True&loc=Local
   
   # Docker MySQL Configuration
   MYSQL_ROOT_PASSWORD=root
   MYSQL_DATABASE=todo-mis
   MYSQL_USER=todo_user
   MYSQL_PASSWORD=todo_password
   ```

3. **Start the application**
   ```bash
   docker compose up -d
   ```

4. **Verify the application is running**
   ```bash
   # Check container status
   docker compose ps
   
   # Check health endpoint
   curl http://localhost:8080/health
   ```

5. **View logs**
   ```bash
   # View all logs
   docker compose logs -f
   
   # View only app logs
   docker compose logs -f app
   ```

6. **Stop the application**
   ```bash
   docker compose down
   
   # Remove volumes (data will be lost)
   docker compose down -v
   ```

### Running Locally

1. **Start MySQL**

   Make sure MySQL 8.0 is running on `localhost:3306`

2. **Create database**
   ```sql
   CREATE DATABASE IF NOT EXISTS `todo-mis` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```

3. **Run database migrations**

   The application will automatically execute `sql/init.sql` on first connection, or you can run it manually:
   ```bash
   mysql -u root -p todo-mis < sql/init.sql
   ```

4. **Install dependencies**
   ```bash
   go mod download
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`

## Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test ./... -cover
```

### Run Tests for Specific Package

```bash
# Test application layer
go test ./application/...

# Test persistence layer
go test ./adapter/driven/persistence/...
```

### Run Tests with Verbose Output

```bash
go test -v ./...
```

### Generate Mock Files

This project uses GoMock for generating mocks:

```bash
# Generate mocks for repository interfaces
mockgen -source=domain/todo/todo_repository.go -destination=domain/todo/mock/mock_todo_repository.go -package=mock
```

## Building

### Build Binary Locally

```bash
# Build for current platform
go build -o bin/todo-mis .

# Build for Linux (useful for Docker)
CGO_ENABLED=0 GOOS=linux go build -o bin/todo-mis .

# Run the built binary
./bin/todo-mis
```

### Build Docker Image

```bash
# Build the image
docker build -t todo-mis-app .

# Run the container
docker run -p 8080:8080 \
  -e MYSQL_DSN="user:password@tcp(mysql-host:3306)/todo-mis?charset=utf8mb4&parseTime=True&loc=Local" \
  todo-mis-app
```

### Build with Docker Compose

```bash
# Rebuild and start
docker compose up -d --build
```

## API Documentation

### Base URL

```
http://localhost:8080
```

### Authentication

Most endpoints require JWT authentication. Include the token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

### Endpoints

#### Health Check

**Get Application Health Status**

```http
GET /health
```

**Response**
- Status: `200 OK`
- Body: Empty (or JSON depending on implementation)

---

#### TODO Items

**Create TODO Item**

```http
POST /api/v1/todo-items
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread"
}
```

**Response**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1
  }
}
```

---

**List TODO Items**

```http
GET /api/v1/todo-items
Authorization: Bearer <token>
```

**Response**
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Milk, eggs, bread",
      "status": 0,
      "user_id": 123,
      "created_at": "2026-01-29T10:00:00Z",
      "updated_at": "2026-01-29T10:00:00Z"
    }
  ]
}
```

---

**Delete TODO Item**

```http
DELETE /api/v1/todo-items/:id
Authorization: Bearer <token>
```

**Response**
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

---

**Complete TODO Item**

```http
PATCH /api/v1/todo-items/:id/complete
Authorization: Bearer <token>
```

**Response**
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

---

#### OAuth Authentication

**Get OAuth Authorization URL**

```http
GET /api/v1/auth/url?provider=google
```

**Query Parameters**
- `provider`: OAuth provider name (`google`, `github`, `facebook`)

**Response**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "redirect_url": "https://accounts.google.com/o/oauth2/v2/auth?..."
  }
}
```

---

**Exchange Authorization Code for Token**

```http
GET /api/v1/auth/token?code=<authorization-code>&provider=google
```

**Query Parameters**
- `code`: Authorization code from OAuth provider
- `provider`: OAuth provider name

**Response**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 123,
      "name": "John Doe"
    }
  }
}
```

### Error Response Format

All errors follow this format:

```json
{
  "code": 400,
  "message": "error description",
  "data": null
}
```

**Common Error Codes**
- `200`: Success
- `400`: Bad Request
- `401`: Unauthorized
- `404`: Not Found
- `500`: Internal Server Error

### cURL Examples

**Create a TODO item**
```bash
curl -X POST http://localhost:8080/api/v1/todo-items \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries",
    "description": "Milk, eggs, bread"
  }'
```

**List TODO items**
```bash
curl -X GET http://localhost:8080/api/v1/todo-items \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Delete a TODO item**
```bash
curl -X DELETE http://localhost:8080/api/v1/todo-items/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Complete a TODO item**
```bash
curl -X PATCH http://localhost:8080/api/v1/todo-items/1/complete \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Project Structure

```
TODO-MIS/
├── adapter/
│   ├── driven/
│   │   ├── auth/              # OAuth provider implementations
│   │   │   ├── auth_provider_factory.go
│   │   │   ├── google_provider.go
│   │   │   ├── github_provider.go
│   │   │   └── facebook_provider.go
│   │   └── persistence/       # Database implementations
│   │       ├── mysql_repository.go
│   │       └── todo_item.go
│   └── driving/
│       └── api/               # HTTP handlers
│           ├── dto/           # Data Transfer Objects
│           │   └── todo.go
│           ├── todo.go
│           └── auth.go
├── application/               # Application services
│   ├── todo.go
│   ├── todo_test.go
│   └── auth.go
├── common/
│   ├── const/                # Constants
│   │   ├── business.go
│   │   └── err.go
│   ├── middware/             # Middleware
│   │   └── auth.go
│   └── util/                 # Utilities
│       └── response.go
├── domain/
│   ├── auth/                 # Auth domain
│   │   ├── auth_service.go
│   │   └── mock/
│   ├── todo/                 # TODO domain
│   │   ├── entity/
│   │   │   └── todo_item.go
│   │   ├── todo_service.go
│   │   ├── todo_repository.go
│   │   └── mock/
│   │       └── mock_todo_repository.go
├── server/                   # Server setup
│   ├── resource.go          # Resource providers (DB, Logger, Gin)
│   └── router.go            # Route registration
├── sql/
│   └── init.sql             # Database initialization
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile              # Docker image definition
├── .env                    # Environment variables
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── main.go                 # Application entry point
└── README.md               # This file
```

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_ENV` | Environment (dev/prod) | `dev` | No |
| `PORT` | HTTP server port | `8080` | No |
| `MYSQL_DSN` | MySQL connection string | - | Yes |
| `MYSQL_ROOT_PASSWORD` | MySQL root password (Docker) | - | Yes (Docker) |
| `MYSQL_DATABASE` | MySQL database name (Docker) | - | Yes (Docker) |
| `MYSQL_USER` | MySQL user (Docker) | - | Yes (Docker) |
| `MYSQL_PASSWORD` | MySQL password (Docker) | - | Yes (Docker) |

### MySQL DSN Format

```
username:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local
```

**Examples:**

- Local: `root:root@tcp(localhost:3306)/todo-mis?charset=utf8mb4&parseTime=True&loc=Local`
- Docker: `todo_user:todo_password@tcp(mysql:3306)/todo-mis?charset=utf8mb4&parseTime=True&loc=Local`

## Development

### Code Generation

**Generate mocks:**
```bash
# For repository interfaces
go generate ./...
```

### Database Migrations

Migrations are automatically applied from `sql/init.sql` when the database is first initialized.

To manually apply migrations:
```bash
mysql -u root -p todo-mis < sql/init.sql
```

### Adding New Features

1. **Define domain entities** in `domain/entity/`
2. **Define repository interface** in `domain/`
3. **Implement repository** in `adapter/driven/persistence/`
4. **Define service interface** in `domain/`
5. **Implement application service** in `application/`
6. **Create DTOs** in `adapter/driving/api/dto/`
7. **Implement HTTP handlers** in `adapter/driving/api/`
8. **Register routes** in `server/router.go`
9. **Wire dependencies** in `main.go` using Fx

### Middleware

All middleware is defined in `common/middware/`:
- `AuthMiddleware`: JWT authentication and authorization

To add middleware to specific routes:
```go
apiGroup.Use(middware.AuthMiddleware())
```

### Logging

The application uses Zap for structured logging:

```go
logger.Info("message",
    zap.String("key", "value"),
    zap.Int("id", 123),
)
```

## Troubleshooting

### Docker Issues

**Problem**: `connection refused` when accessing the app

**Solution**: Ensure the server is binding to `0.0.0.0:8080` (all interfaces), not `localhost:8080`

---

**Problem**: MySQL connection refused

**Solution**: 
1. Check if MySQL container is running: `docker compose ps`
2. Verify environment variables in `.env`
3. Check logs: `docker compose logs mysql`

---

**Problem**: Container keeps restarting

**Solution**: Check app logs for errors: `docker compose logs app`

### Database Issues

**Problem**: `Access denied for user 'todo_user'`

**Solution**: 
1. Verify `MYSQL_USER` and `MYSQL_PASSWORD` in `.env`
2. Recreate containers with clean volumes: `docker compose down -v && docker compose up -d`

---

**Problem**: Database initialization failed

**Solution**: Check `sql/init.sql` for syntax errors and MySQL logs


