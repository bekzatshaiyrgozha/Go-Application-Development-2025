.PHONY: sqlc db-migrate

# Генерация Go кода из SQL
sqlc:
	sqlc generate

# Запуск миграций
db-migrate:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up