generate-proto:
	mkdir -p ./pkg/api
	protoc \
		--go_out=./pkg \
		--go_opt=paths=source_relative \
		--go-grpc_out=./pkg \
		--go-grpc_opt=paths=source_relative \
		./api/lodas.proto
	mkdir -p ./pkg/code
	protoc --go_out=./pkg/code --go-grpc_out=./pkg/code  ./api/lodas_code.proto

generate-ts:
	mkdir -p ./pkg/ts
	protoc -I=./api --js_out=import_style=commonjs:./pkg/ts --grpc-web_out=import_style=typescript,mode=grpcwebtext:./pkg/ts ./api/lodas.proto
	protoc -I=./api --js_out=import_style=commonjs:./pkg/ts --grpc-web_out=import_style=typescript,mode=grpcwebtext:./pkg/ts ./api/lodas_code.proto

generate-ent:
	mkdir -p ./pkg/ent
	cp -R ./schema ./pkg/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema 