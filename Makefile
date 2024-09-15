proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest			

	protoc --go-grpc_out=. --go_out=. ./protobuf/**/**.proto


