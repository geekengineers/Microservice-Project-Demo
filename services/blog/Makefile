dev:
	export GRPC_GO_LOG_VERBOSITY_LEVEL=99
	export GRPC_GO_LOG_SEVERITY_LEVEL=info
	go run ./cmd/grpc/main.go

integration_test:
	go test -cover -coverpkg ./internal/core/services/... ./tests/integration/...

genrate_jwt_secret: 
	ruby ./config/generate_jwt_secret.rb