compileAll: account apiGateway

account:
	cd service/account && make compile

apiGateway:
	cd service/api_gateway && make compile
