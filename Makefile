DB_URL=postgres://postgres:postgres@db:5432/postgres?sslmode=disable

migrateup:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose drop

test-unit:
	go test -v -cover -count 1 -run ^TestUnit ./...

test-integration:
	go test -v -cover -p 1 -count 1 -run ^TestIntegration ./...

test-clean:
	go clean --testcache

.PHONY:  migrateup migratedown test test-clean
