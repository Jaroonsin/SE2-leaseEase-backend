.PHONY: run docker swag

# 'swag' target to initialize Swagger
swag:
	swag init -g cmd/main.go -o cmd/docs/v2  --ot yaml
# 'run' target to run the Go application and the Docker containers
docker:
	docker-compose up -d
	air
	go run cmd/main.go

# 'run2' target to just run the Go application
run:
	air
	go run cmd/main.go

rm_docker:
	docker-compose down
	docker volume rm se2-leaseease-backend_postgres_data
	docker volume rm se2-leaseease-backend_pgadmin_data 

