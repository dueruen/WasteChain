compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/QR.proto --go_out=plugins=grpc:gen

build: compile
	cd cmd && GOOS=linux go build main.go
	docker build -t qr-service .

include ../../base.mk
