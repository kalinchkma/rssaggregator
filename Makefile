start:
	@go run ./*.go

docker-build:
	@docker compose build

docker-down:
	@docker compose down

docker-up:
	@docker compose up -d