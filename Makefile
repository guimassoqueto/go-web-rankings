db:
	docker compose up -d postgres

run-classic:
	go run cmd/classic/main.go

run-pgx:
	go run cmd/pgx/main.go

run-gorm:
	go run cmd/gorm/main.go

or:
	open https://github.com/guimassoqueto/go-web-rankings