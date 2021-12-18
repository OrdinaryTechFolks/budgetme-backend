setup:
	go mod vendor

run:
	go build -v -o bin/app-grpc cmd/app-grpc/*.go && ./bin/app-grpc --debug