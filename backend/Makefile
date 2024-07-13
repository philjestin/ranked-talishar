test:
		go test -v --cover ./...

postgres:
		docker-compose up

serve:
		air

sqlc:
		sqlc generate

migrateup:
		migrate -path db/migration -database  "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database  "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose down

.PHONY: postgres test serve sqlc migrateup migratedown