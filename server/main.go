package main

//go:generate bash -c "protoc --proto_path=../protobuf --go_out=generated/ --go-grpc_out=generated/ $(find ../protobuf -type f -name *.proto)"

func main() {
}
