
.PHONY: gen.protoc
gen.protoc:
	@protoc --proto_path=$(APP_ROOT)/pkg/api \
	       --proto_path=$(APP_ROOT)/third_party \
 	       --go_out=paths=source_relative:$(APP_ROOT)/pkg/api \
 	       --go-http_out=paths=source_relative:$(APP_ROOT)/pkg/api \
 	       --go-grpc_out=paths=source_relative:$(APP_ROOT)/pkg/api \
		   --go-errors_out=paths=source_relative:$(APP_ROOT)/pkg/api \
	       $(shell find $(APP_ROOT)/pkg/api -name *.proto)
	@protoc --proto_path=$(APP_ROOT)/internal \
	       --proto_path=$(APP_ROOT)/third_party \
 	       --go_out=paths=source_relative:$(APP_ROOT)/internal \
	       $(shell find $(APP_ROOT)/internal -name *.proto)

