# ranked-talishar

## Overview and Setup

sqlc: Is a library that is used to generate Golang code that has type safe interfaces from SQL queries.

air: Is a live reloading utility that is used to observe changes in Go applications. Run the air command in the root of your project and it will take care of listening to changing events in your code without the need to restart the application.

pq: It's GO Postgres driver used for GO database/sql. You have to import this driver for you to use the database/sql if you are using the Postgresql engine.

viper: Is a configuration package solution used to read configurations of all types ranging from environment variables, buffer files, command line flags to files written in JSON, YAML and TOML to name just but a few.

gin: Is a web framework written in Go and it comes with a lot of prebuilt in features ranging from routes grouping, JSON validation, middleware support and error management.

migrate: This library helps with database migration either through incremental changes in the schema or reversal changes made in the database schema.

## Reminders
Running docker `docker-compose up`

Running golang: `air`

Creating up/down migration files: `migrate create -ext sql -dir db/migration -seq name_of_schema`

SQL Queries: `sqlc generate`