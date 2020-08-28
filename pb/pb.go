package main

//go:generate protoc --go_out=../pb ../pb/protocolVersion.proto -I=../pb
//go:generate protoc --go_out=../pb ../pb/v1/region_API.proto  -I=../pb
//go:generate protoc --go_out=../pb ../pb/v1/locator_API.proto  -I=../pb
//go:generate protoc --go_out=../pb ../pb/v1/function_API.proto  -I=../pb
//go:generate protoc --go_out=../pb ../pb/v1/connection_API.proto  -I=../pb
//go:generate protoc --go_out=../pb ../pb/v1/clientProtocol.proto  -I=../pb
//go:generate protoc --go_out=../pb ../pb/v1/basicTypes.proto  -I=../pb

func main() {

}
