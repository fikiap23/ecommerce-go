# up:
# 	nodemon --watch './**/*.go' --signal SIGINT --exec APP_ENV=dev 'go' run main.go

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


