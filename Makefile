setup:
	go mod vendor

build-proto:
	@protoc -I grpc/budgetme/proto grpc/budgetme/proto/*.proto --gofast_out=plugins=grpc:grpc/budgetme/proto

build-run:
	go build -v -o bin/app-grpc cmd/app-grpc/*.go && ./bin/app-grpc --debug

start-dev: 
	reflex -r \.go -s make build-run