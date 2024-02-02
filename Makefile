build: generate
	go build -o ./dist/main ./cmd/server/main.go

dev:
	air

generate:
	templ generate
	./tailwindcss -o ./static/vendor/tailwind.css

run: build
	./dist/main

build-migrate:
	go build -o ./dist/migrate ./cmd/migrate/main.go

migrate: build-migrate
	./dist/migrate
