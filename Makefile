.SILENT:

.PHONY: lint
lint:
	golangci-lint run

create-migration:
	migrate create -ext sql -dir schema/ -seq $(NAME)

migrate-up:
	migrate -path ./schema -database "postgresql://postgres:password@localhost:5437/postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./schema -database "postgresql://postgres:password@0.0.0.0:5437/postgres?sslmode=disable"  -verbose down

migrate-fix: 
	migrate -path ./schema -database "postgresql://postgres:password@0.0.0.0:5437/postgres?sslmode=disable"  force $(VERSION)

clean-migration:
	del /Q schema\$(FILENAME)