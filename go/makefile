# docker
dev:
	docker compose -f dev-docker-compose.yml up --build

dev-down:
	docker compose -f dev-docker-compose.yml down

test:
	docker compose -f test-docker-compose.yml up --build

test-down:
	docker compose -f test-docker-compose.yml down

# linter
lint:
	golangci-lint run

# ent generate
ent-generate:
	go generate ./ent