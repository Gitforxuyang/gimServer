proto:
	protoc --go_out=plugins=grpc:./proto/ -I=./proto/ im.proto

gim:
	mkdir proto/gim | true
	protoc --proto_path=proto --go_out=proto/gim gim.proto

.PHONY: proto