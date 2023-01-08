# Microservice Architecture Playgorund

## Migration

### Run migration

To run database migration, first install [golang-migrate](https://github.com/golang-migrate/migrate) and then run the following command:

```bash

migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

```

Other useful migration commands can be found in [golang-migrate](https://github.com/golang-migrate/migrate) github page.

### Update migration file according to gorm model

To update migration file according to gorm model, run the migration main file with "name" flag:

```bash
go run cmd/migration/main.go -name=add-field-to-user
```

This will create a new migration file in `migrations` folder. Then you can run the migration command above to update the database.
