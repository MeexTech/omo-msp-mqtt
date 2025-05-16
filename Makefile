.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/mqtt/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/mqtt/plug.proto
