.PHONY: up
up:
	docker-compose up -d

.PHONY: build
build:
	docker-compose build

.PHONY: exec
exec:
	docker-compose exec pypy /bin/bash

.PHONY: stop
stop:
	docker-compose stop
