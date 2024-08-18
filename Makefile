# Load the environment variables from .env
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Target to log in to the PostgreSQL database
.PHONY: db-login
db/login:
	docker compose exec db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

docker/build:
	docker build -t fam-chat .

docker/run:
	docker run -p 8080:8080 fam-chat

scripts/create_migration:
	@read -p "And you traveler, what's your migration called: " migration && ./scripts/create_migration.sh $${migration}

