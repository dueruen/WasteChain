compileAll: account apiGateway signature qr authentication blockchain

account:
	cd service/account && make compile

apiGateway:
	cd service/api_gateway && make compile

signature:
	cd service/signature && make compile

qr:
	cd service/qr && make compile

authentication:
	cd service/authentication && make compile

blockchain:
	cd service/blockchain && make compile
