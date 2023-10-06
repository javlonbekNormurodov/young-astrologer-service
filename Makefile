build:
	docker build -t young-astrologer-service .

run:
	docker-compose up -d

clean:
	docker-compose down

worker:
	go run worker/worker.go

api:
	go run api/api.go
