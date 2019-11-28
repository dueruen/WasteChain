compile: clean protoc graphql

protoc:
	protoc -I ../.. ../../proto/Account.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Signature.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/QR.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Shipment.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Authentication.proto --go_out=plugins=grpc:gen


graphql:
	go run cmd/gqlgen.go

include ../../base.mk
