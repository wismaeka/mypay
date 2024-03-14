migrate-local:
	docker compose -f dev.docker-compose.yaml up migrate

db-local:
	docker compose -f dev.docker-compose.yaml up -d database

migrate-down:
	docker compose -f dev.docker-compose.yaml run --rm migrate down ${step}

migrate-force:
	docker compose -f dev.docker-compose.yaml run --rm migrate force ${v}

tear:
	docker compose -f dev.docker-compose.yaml down

sleep-one-sec:
	sleep 1;

setup: db-local sleep-one-sec migrate-local

run:
	go run app/main.go

run-local: setup
	go run app/main.go
