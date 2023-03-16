generate_grpc_code:
	protoc --go_out=grpsServer --go_opt=paths=source_relative --go-grpc_out=grpsServer --go-grpc_opt=paths=source_relative grpcServer.proto
