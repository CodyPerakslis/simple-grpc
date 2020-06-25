
.PHONY: protobuf
protobuf:
	protoc -I messenger/ messenger/messenger.proto --go_out=plugins=grpc:messenger --go_opt=paths=source_relative