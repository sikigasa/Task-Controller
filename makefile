.PHONY: genswag genproto

genswag:
	protoc -I . --openapiv2_out ./docs --openapiv2_opt allow_merge=true,disable_default_errors=true $(file)

genproto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/v1/*.proto