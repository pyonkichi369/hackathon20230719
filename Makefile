up:
	docker compose up -d
build:
	docker compose build
down:
	docker compose down --remove-orphans
restart:
	@make down
	@make up
ps:
	docker compose ps
app:
	docker compose exec app /bin/sh
