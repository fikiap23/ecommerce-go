# up:
# 	nodemon --watch './**/*.go' --signal SIGINT --exec APP_ENV=dev 'go' run main.go

# 🟢 Development commands
up:
	docker compose up -d
	docker compose logs -f backend-go

up-build:
	docker compose up --build

exec:
	docker exec -it backend-go sh

down:
	docker compose down

restart:
	docker restart backend-go
	docker compose logs -f backend-go

logs:
	docker compose logs -f backend-go

# 🧪 Test environment commands
test-up:
	docker compose -f docker-compose.test.yml up -d
	docker compose -f docker-compose.test.yml logs -f backend-go-test

test-build:
	docker compose -f docker-compose.test.yml up --build

test-down:
	docker compose -f docker-compose.test.yml down

test-logs:
	docker compose -f docker-compose.test.yml logs -f backend-go-test

test-exec:
	docker exec -it backend-go-test sh

test-run:
	docker compose -f docker-compose.test.yml run --rm backend-go-test go test -v ./test/e2e
	docker compose -f docker-compose.test.yml down

test-all: test-down test-build

