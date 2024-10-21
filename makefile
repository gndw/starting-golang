build-app:
	@echo "Building App..."
	@go build -o bin/main cmd/main.go

run-app: build-app
	@echo "Running App..."
	@bin/main

hit-test:
	@curl -i -X POST localhost:8080/test -d '{"user_id":100}' 

test:
	@echo "Unit Testing..."
	@go test -coverprofile ./coverage.out ./...
