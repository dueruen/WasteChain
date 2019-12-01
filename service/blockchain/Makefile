compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Blockchain.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Signature.proto --go_out=plugins=grpc:gen

build: compile
	cd cmd && go build main.go
	docker build -t blockchain-service .

include ../../base.mk
