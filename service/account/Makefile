compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Account.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Authentication.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Signature.proto --go_out=plugins=grpc:gen

build: compile
	cd cmd && go build main.go
	docker build -t account-service .

include ../../base.mk
