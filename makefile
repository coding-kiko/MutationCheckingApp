compile-arm:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/main
build:
	docker-compose up -d --build
down:
	docker-compose down
