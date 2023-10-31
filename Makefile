start: build run
build:
	docker build -t coinwallet .
run:
	docker run -p 8090:8090 -v ./pb_data:/app/pb_data coinwallet
local:
	go run main.go serve --http=127.0.0.1:8090
dev:
	air
