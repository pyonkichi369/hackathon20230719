up:
	docker compose up -d
build:
	docker compose build
down:
	docker compose down --remove-orphans
restart:
	@make down
	@make up
app:
	docker compose exec app bash
