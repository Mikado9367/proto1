# Settlement Broker service

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative secpos.proto


Broker service est à la fois client et server gRPC.

Du coup, pour la partie server, il faut s'inspirer de log-service. 
Et pour la partie client, il faut s'inspirer de broker-service.

Maintenant, la question est : est ce que cela a du sens car finalement,
le broker n'est qu'un passe plat, avec des routes.

-   flux cash balances
-   flux security positions
