# Gemini Context: starting-golang

This document provides a technical overview of the `starting-golang` repository to accelerate future analysis and development sessions.

## 🏗️ Architecture & Design Patterns
The project follows a **layered architecture** with strict separation of concerns, ensuring high testability and maintainability.

- **Dependency Layer (`internals/dependencies`):** Wraps external libraries and standard packages (e.g., `os`, `godotenv`) into interfaces for improved testability and mocking.
- **Service Layer (`internals/services`):** Utility-focused components with single responsibilities. Current services include:
    - `env`: Manages configuration from `.env` and system environment variables.
    - `log`: Provides a structured logging interface.
    - `httpmiddlewarelog`: HTTP middleware for logging requests and responses.
    - `httpserver`: Custom HTTP server implementation using standard library `http.ServeMux`.
- **Handler Layer (`internals/handlers`):** Entry point for external requests. Responsible for parsing request bodies (JSON) and handling HTTP-specific logic.
- **Usecase Layer (`internals/usecase`):** Contains the core business logic. It is agnostic of the transport layer (HTTP, gRPC, etc.).
- **Repository Layer (`internals/repositories`):** Manages data persistence and external integrations (e.g., In-memory DB).

### Dependency Injection
All layers are wired together in `internals/resources/resource.go`. Each layer typically defines an **Interface** in its own package and provides a concrete **Implementation**.

## 🚀 Key Entry Points
- **Main Entry:** `cmd/main.go` - Initializes resources and starts the HTTP server.
- **Wiring:** `internals/resources/resource.go` - The `Init` function where all services and data layers are instantiated.
- **Server:** `internals/services/httpserver/implementation.go` - Implements the HTTP server using `http.ServeMux`. (Start() method)

## 🛠️ Tech Stack & Tooling
- **Language:** Go 1.26.0
- **HTTP Framework:** Standard library (`net/http`) with a custom `RegisterEndpoint` abstraction.
- **Testing:** `github.com/stretchr/testify` for assertions and `mockery` for generating mocks. Generate mocks only when all files are saved without errors.
- **Automation:** `makefile` handles building, running, testing, and hitting endpoints.

## 📋 Common Workflows

### Running the Application
```bash
make run-app
```
Starts the server on `:5548`.

### Verifying the Server
```bash
make hit-test
```
Sends a POST request to the `/test` endpoint with a sample payload.

### Testing
```bash
make test
```
Runs all unit tests with coverage reporting.

### Adding a New Endpoint
1. Define the request/response models in `internals/models`.
2. Implement business logic in a new or existing **Usecase**.
3. Create a **Handler** that uses the Usecase.
4. Register the handler's endpoint in `internals/resources/resource.go` using `httpServerService.RegisterEndpoint`.

## 📂 Directory Map
- `cmd/`: Application entry point.
- `internals/dependencies/`: Testable wrappers for external dependencies.
- `internals/handlers/`: HTTP request handlers.
- `internals/usecase/`: Business logic implementations.
- `internals/repositories/`: Data access layer.
- `internals/services/`: Cross-cutting concerns (logging, server, middleware).
- `internals/models/`: Shared data structures.
- `internals/constants/`: Shared constants and types.
