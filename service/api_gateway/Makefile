compile: clean protoc graphql

protoc:
	protoc -I ../.. ../../proto/Account.proto --go_out=plugins=grpc:gen

graphql:
	go run cmd/gqlgen.go

build: compile
	cd cmd && go build main.go
	docker build -t api-service .

include ../../base.mk
