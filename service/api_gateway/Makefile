compile: clean protoc graphql

protoc:
	protoc -I ../.. ../../proto/Account.proto --go_out=plugins=grpc:gen

graphql:
	go run cmd/gqlgen.go

include ../../base.mk
