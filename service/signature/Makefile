compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/Signature.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/QR.proto --go_out=plugins=grpc:gen
	protoc -I ../.. ../../proto/Blockchain.proto --go_out=plugins=grpc:gen

build: compile
	cd cmd && GOOS=linux go build main.go
	docker build -t signature-service .

include ../../base.mk
