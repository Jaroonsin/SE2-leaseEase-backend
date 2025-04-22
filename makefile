.PHONY: run docker swag
.PHONY: test

# 'swag' target to initialize Swagger
swag:
	swag init -g cmd/main.go -o cmd/docs/v2  --ot yaml
# 'run' target to run the Go application and the Docker containers
docker:
	docker-compose build --no-cache
	docker-compose up -d
	air
	export API_HOST="staging.example.com"
	export API_PORT="8080"
	go run cmd/main.go

# 'run2' target to just run the Go application
run:
	air
	export API_HOST="staging.example.com"
	export API_PORT="8080"
	go run cmd/main.go

rm_docker:
	docker-compose down -v
	docker volume rm leaseease_postgres_data
	docker volume rm leaseease_pgadmin_data

test:
	go test -v ./test/main_test.go

