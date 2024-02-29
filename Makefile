.SILENT:

.PHONY: lint
lint:
	golangci-lint run

create-migration:
	migrate create -ext sql -dir schema/ -seq $(NAME)

.PHONY: migrate-up
migrate-up:
	migrate -path ./schema -database "postgresql://postgres:password@postgres:5437/postgres?sslmode=disable" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path ./schema -database "postgresql://postgres:password@postgres:5437/postgres?sslmode=disable"  -verbose down

.PHONY: migrate-fix
migrate-fix: Ц
	migrate -path ./schema -database "postgresql://postgres:password@postgres:5437/postgres?sslmode=disable"  force $(VERSION)

.PHONY: clean-migration
clean-migration:
	rm -f schema/$(FILENAME)
