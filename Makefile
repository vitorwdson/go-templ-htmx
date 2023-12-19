build:
	go build -o ./dist/main ./cmd/main.go

dev:
	air

generate:
	templ generate
	./tailwindcss -o ./static/vendor/tailwind.css

run:
	./dist/main
