package proto

//go:generate ${PROTOC} -I${GOPATH}/src/github.com/google/protobuf/src -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. chat.proto
