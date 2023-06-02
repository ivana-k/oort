protoc --proto_path=proto/common \
	--go_out=proto/common \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/common \
	--go-grpc_opt=paths=source_relative \
	 common.proto
protoc -I=proto/common \
	--proto_path=proto/checkerpb \
	--go_out=proto/checkerpb \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/checkerpb \
	--go-grpc_opt=paths=source_relative \
	checker.proto
protoc -I=proto/common \
	--proto_path=proto/syncerpb \
	--go_out=proto/syncerpb \
	--go_opt=paths=source_relative  \
	--go-grpc_out=proto/syncerpb \
	--go-grpc_opt=paths=source_relative \
	syncer.proto
