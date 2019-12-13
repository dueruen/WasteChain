buildAll: account apiGateway signature qr authentication blockchain shipment

runDev: buildAll
	docker-compose up

down:
	docker-compose down

account:
	cd service/account && make build

apiGateway:
	cd service/api_gateway && make build

qr:
	cd service/qr && make build

shipment:
	cd service/shipment && make build

signature:
	cd service/signature && make build

authentication:
	cd service/authentication && make build

blockchain:
	cd service/blockchain && make build
