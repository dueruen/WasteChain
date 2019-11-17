compileAll: account apiGateway qr shipment

account:
	cd service/account && make compile

apiGateway:
	cd service/api_gateway && make compile

qr:
	cd service/qr && make compile

shipment:
	cd service/shipment && make compile
