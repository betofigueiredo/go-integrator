ARGS = $(filter-out $@,$(MAKECMDGOALS))
MAKEFLAGS += --silent

b-api:
	docker compose build api

b-integrator:
	docker compose build integrator

up:
	docker compose up

down:
	docker compose down
