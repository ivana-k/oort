protoc --proto_path=pkg/proto \
	--go_out=pkg/proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=pkg/proto \
	--go-grpc_opt=paths=source_relative \
	 model.proto
protoc -I=pkg/proto \
	--proto_path=pkg/proto \
	--go_out=pkg/proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=pkg/proto \
	--go-grpc_opt=paths=source_relative \
	evaluator.proto
protoc -I=pkg/proto \
	--proto_path=pkg/proto \
	--go_out=pkg/proto \
	--go_opt=paths=source_relative  \
	--go-grpc_out=pkg/proto \
	--go-grpc_opt=paths=source_relative \
	administrator.proto
