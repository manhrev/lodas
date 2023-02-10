generate-proto:
	mkdir -p ./pkg/api
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ./api/auth.proto

generate-code:
	mkdir -p ./pkg/code
	protoc --go_out=./pkg/code --go-grpc_out=./pkg/code  ./api/auth_code.proto

generate-ts:
	mkdir -p ./npm/code
	protoc -I=./api --js_out=import_style=commonjs:./npm/code --grpc-web_out=import_style=typescript,mode=grpcwebtext:./npm/code ./api/auth.proto
	protoc -I=./api --js_out=import_style=commonjs:./npm/code --grpc-web_out=import_style=typescript,mode=grpcwebtext:./npm/code ./api/auth_code.proto

generate-cache:
	mkdir -p ./pkg/data
	protoc --go_out=./pkg/data --go-grpc_out=./pkg/data  ./api/auth_cache.proto

generate-ent:
	mkdir -p ./pkg/ent
	cp -R ./schema ./pkg/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema

#rm -r ./pkg/ent/schema
generate:
	make generate-proto generate-code generate-ent generate-cache

run:
	go run ./cmd/.
