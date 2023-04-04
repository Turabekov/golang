

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://postgres:admin@123@localhost:5432/catalog?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://postgres:admin@123@localhost:5432/catalog?sslmode=disable' down
