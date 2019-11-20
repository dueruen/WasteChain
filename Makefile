compileAll: account apiGateway signature qr

account:
	cd service/account && make compile

apiGateway:
	cd service/api_gateway && make compile

signature:
	cd service/signature && make compile

qr:
	cd service/qr && make compile
