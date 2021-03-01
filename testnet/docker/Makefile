all: mynode rest-server

mynode:
	docker build --tag shareledger/mynode shareledgernode

rest-server:
	docker build --tag shareledger/rest-server rest-server

.PHONY: all rest-server mynode

