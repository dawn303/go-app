# ==============================================================================
#  Makefile helper functions for tools
#
# Specify tools category.
KRATOS_INIT_TOOLS = kratos protoc-gen-go protoc-gen-go-grpc protoc-gen-go-http protoc-gen-openapi wire

.PHONY: _install.kratos
_install.kratos: 
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
