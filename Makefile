proto:
	protoc --go_out=plugins=grpc:./proto/ -I=./proto/ im.proto

.PHONY: proto