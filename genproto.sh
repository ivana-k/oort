protoc --proto_path=proto/common \
	--go_out=proto/common \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/common \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out proto/common \
    	--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
	 common.proto
protoc -I=proto/common \
	--proto_path=proto/checkerpb \
	--go_out=proto/checkerpb \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/checkerpb \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out proto/checkerpb \
    	--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
	checker.proto
protoc -I=proto/common \
	--proto_path=proto/syncerpb \
	--go_out=proto/syncerpb \
	--go_opt=paths=source_relative  \
	--go-grpc_out=proto/syncerpb \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out proto/syncerpb \
    	--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
	syncer.proto
