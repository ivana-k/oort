protoc --proto_path=./ \
	--go_out=../ \
	--go_opt=paths=source_relative \
	--go-grpc_out=../ \
	--go-grpc_opt=paths=source_relative \
	 model.proto
protoc --proto_path=./ \
	--go_out=../ \
	--go_opt=paths=source_relative \
	--go-grpc_out=../ \
	--go-grpc_opt=paths=source_relative \
	 administrator_async.proto
protoc -I=. \
	--proto_path=./ \
	--go_out=../ \
	--go_opt=paths=source_relative \
	--go-grpc_out=../ \
	--go-grpc_opt=paths=source_relative \
	evaluator.proto
protoc -I=. \
	--proto_path=./ \
	--go_out=../ \
	--go_opt=paths=source_relative  \
	--go-grpc_out=../ \
	--go-grpc_opt=paths=source_relative \
	administrator.proto
