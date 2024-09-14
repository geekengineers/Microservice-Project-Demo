dev:
	export GRPC_GO_LOG_VERBOSITY_LEVEL=99
	export GRPC_GO_LOG_SEVERITY_LEVEL=info
	go run ./cmd/grpc/main.go

integration_test:
	go test -cover -coverpkg ./internal/core/services/... ./tests/integration/...

proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest	
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest			
	protoc --go-grpc_out=. --go_out=. ./protobuf/auth/auth.proto	

genrate_jwt_secret: 
	ruby ./config/generate_jwt_secret.rb