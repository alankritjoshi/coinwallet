start: build run
build:
	docker build -t coinwallet .
run:
	docker run -p 8080:8080 -v ./pb_data:/app/pb_data -v ./pb_migrations:/app/pb_migrations coinwallet
local:
	go run app/main.go serve --http=127.0.0.1:8080
dev:
	air
