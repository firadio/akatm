@echo off
goctl rpc protoc service.proto --proto_path=. -m --go_out=../pb --go-grpc_out=../pb --zrpc_out=.. --style go_zero
pause
