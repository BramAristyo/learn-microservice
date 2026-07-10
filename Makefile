gen:
	protoc \
		-I=proto \
		--go_out=. \
		--go_opt=module=github.com/bramAristyo/learn-microservice \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/bramAristyo/learn-microservice \
		proto/user/user.proto

gen-example:
	protoc \
		-I=. \
		--go_out=. \
		--go_opt=module=github.com/bramAristyo/learn-microservice \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/bramAristyo/learn-microservice \
		grpc-playground/example.proto

