.SILENT:

.PHONY: lint
lint:
	golangci-lint run

create-migration:
	migrate create -ext sql -dir schema/ -seq $(NAME)

migrate-up:
	migrate -path schema/ -database "postgresql://postgres:password@localhost:5432/flightbook?sslmode=disable" -verbose up

migrate-down:
	migrate -path schema/ -database "postgresql://postgres:password@localhost:5432/flightbook?sslmode=disable"  -verbose down

migrate-fix: 
	migrate -path schema/ -database "postgresql://postgres:password@localhost:5432/flightbook?sslmode=disable"  force $(VERSION)

clean-migration:
	del /Q schema\$(FILENAME)