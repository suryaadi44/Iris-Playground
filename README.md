# Microservice Architecture Playgorund

## Migration

To run database migration, first install [golang-migrate](https://github.com/golang-migrate/migrate) and then run the following command:

```bash

migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

```

or run make command:

```bash

make migrate

```

Other useful migration commands can be found in [golang-migrate](https://github.com/golang-migrate/migrate) github page.