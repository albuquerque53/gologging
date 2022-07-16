up:
	@echo "Starting up containers in daemon mode..."
	docker-compose -f build/docker-compose.yml up -d
	@echo "Finished!"

app:
	@echo "Getting into container..."
	docker exec -it gologging_app /bin/bash

run:
	go run ./cmd/main.go

down:
	@echo "Shutting down containers..."
	docker-compose -f build/docker-compose.yml up -d
	@echo "Finished!"
