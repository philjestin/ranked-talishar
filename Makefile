#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up


migrate-down:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1


migrate-create:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)


shell-db:
	docker compose -f ${DOCKER_COMPOSE_FILE} exec postgres psql -U root -d postgres

.PHONY: migrate-up migrate-down migrate-create shell-db