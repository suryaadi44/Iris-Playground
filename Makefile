genproto:
	
	protoc --proto_path=app/api/grpc/proto \
		--go_out=app/api/grpc/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/api/grpc/pb --go-grpc_opt=paths=source_relative \
		app/api/grpc/proto/*.proto 

migrate-up:
	migrate -source file://migrations -database "postgres://root:root@localhost:5432/root?sslmode=disable" up

migrate-down:
	migrate -source file://migrations -database "postgres://root:root@localhost:5432/root?sslmode=disable" down