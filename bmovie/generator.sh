protoc -I/usr/local/include -I/usr/local/Cellar/protobuf/3.15.5/include -I. -I$GOPATH/src \
  --go_out . \
  --go_opt module=github.com/danClauz/bibit/bmovie \
  --go-grpc_out . \
  --go-grpc_opt module=github.com/danClauz/bibit/bmovie \
  src/search.proto

protoc -I/usr/local/include -I/usr/local/Cellar/protobuf/3.15.5/include -I. -I$GOPATH/src \
  --grpc-gateway_out . \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt module=github.com/danClauz/bibit/bmovie \
  src/search.proto

protoc -I/usr/local/include -I/usr/local/Cellar/protobuf/3.15.5/include -I. -I$GOPATH/src \
  --openapiv2_out search/gen \
  src/search.proto