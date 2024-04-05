generate:
	# protoc --proto_path= /*.proto --go_out=. --go-grpc_out=.
	buf generate

run:
	air --build.cmd "go build -o bin/server server/main.go" --build.bin "./bin/server"

ui:
	grpcui --plaintext localhost:8080
