genproto:
	rm -rf app/api/grpc/pb/*.go
	protoc --proto_path=app/api/grpc/proto \
		--go_out=app/api/grpc/pb --go_opt=paths=source_relative \
		--go-grpc_out=app/api/grpc/pb --go-grpc_opt=paths=source_relative \
		app/api/grpc/proto/*.proto
