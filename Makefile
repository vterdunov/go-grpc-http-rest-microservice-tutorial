proto:
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 todo-service.proto
	protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 todo-service.proto

mysql:
	docker run -it --rm -v $(CURDIR)/Create_ToDo_Table.sql:/docker-entrypoint-initdb.d/Create_ToDo_Table.sql -e MYSQL_DATABASE=grpc -e MYSQL_USER=grpc -e MYSQL_PASSWORD=grpc -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306 mysql
