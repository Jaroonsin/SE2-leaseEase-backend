.PHONY: run

run:
	swag init -g cmd/main.go -o cmd/docs
	docker-compose up -d
	go run cmd/main.go