ifneq (,$(wildcard ./.env))
    include .env
    export
endif

CSS_INPUT = assets/src/style.css
CSS_OUTPUT = assets/dist/style.css
PROXY_PORT = 8080
APP_PORT = ${HTTP_PORT}

MIGRATION_DIR = db/migrations
MIGRATION_DRIVER = ${DB_DRIVER}
MIGRATION_DB_STRING=${DB_STRING}


new-migration:
ifndef NAME
	$(error Usage: make new-migration NAME=your_migration_name)
endif
	goose -dir $(MIGRATION_DIR) create $(NAME) sql

migrate-up:
	goose -dir $(MIGRATION_DIR) $(MIGRATION_DRIVER) $(MIGRATION_DB_STRING) up
	sqlc generate

migrate-down:
	goose -dir $(MIGRATION_DIR) $(MIGRATION_DRIVER) $(MIGRATION_DB_STRING) down
	sqlc generate

migrate-reset:
	goose -dir $(MIGRATION_DIR) $(MIGRATION_DRIVER) $(MIGRATION_DB_STRING) reset
	sqlc generate

migrate-status:
	goose -dir $(MIGRATION_DIR) $(MIGRATION_DRIVER) $(MIGRATION_DB_STRING) status


build:
	templ generate 
	npx tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT)
	go build -tags embed -o ./tmp/build


css-watch:
	npx tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch=always

templ-watch:
	templ generate -watch \
		-proxy="http://localhost:$(APP_PORT)" \
		-proxyport="$(PROXY_PORT)" \
		--open-browser=false \
		-cmd="go run ."

dev:
	parallel -j 2 --line-buffer ::: \
		'make css-watch' \
		'make templ-watch'


install-deps:
	npm install
	go mod tidy
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/a-h/templ/cmd/templ@latest
	sudo apt install parallel
