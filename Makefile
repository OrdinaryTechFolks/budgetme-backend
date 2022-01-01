setup:
	go mod vendor

build-proto:
	@protoc -I grpc/budgetme/proto grpc/budgetme/proto/*.proto --gofast_out=plugins=grpc:grpc/budgetme/proto

build:
	go build -v -o bin/app-grpc cmd/app-grpc/*.go

start-dev: 
	reflex -r "\.(go|yaml)" -s -- sh -c "make build && ./bin/app-grpc -config=./files/config/development.yaml"

start-prod:
	./bin/app-grpc -config=./files/config/production.yaml 