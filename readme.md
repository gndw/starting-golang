# Starting Golang
Simple starting golang repositor with existing http server

## Quick Test
- Start the service by `make run-app`

## Test Report
- Run `go mod vendor` and `make test` to run all tests
- Run `go tool cover -html=coverage.out` to check coverage

## Architecture Solution
- Code is separated into different layers and parts
  - Service is a utility part that has single responsibility (ex: httpserver service to start HTTP Server)
  - Handler is a data layer that process incoming HTTP request, reading request body and all http related (ex: http handler, grpc handler, pubsub handler, etc)
  - Usecase is a data layer that handle business logic only
  - Repository is a data layer that handle outgoing access (ex: database, cache, external service, 3rd party, etc)
- Each data layer consist of Interface & Implementation, so that Unit Test can be created easily in each layer
- Every service and data layers is initiated in the `resource` function
