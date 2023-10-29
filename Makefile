start: build run
build:
	docker build -t coinwallet .
run:
	docker run -p 8080:8080 coinwallet
local:
	go run app/main.go serve --http=127.0.0.1:8080
dev:
	air
