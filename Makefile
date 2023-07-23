SHELL=/bin/bash

include .env.dev

.PHONY: db-migrate
db-migrate:
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -source "file://ddl" -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(LOCAL_DB_HOST):$(DB_PORT))/$(DB_DATABASE)" up

.PHONY: gen-api
gen-api:
	mkdir -p ./gen/api
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	oapi-codegen --config config/oapi-codegen/server.yaml ./api/openapi.yaml

.PHONY: up
up:
	docker-compose up

.PHONY: down
down:
	docker-compose down

.PHONY: seed # seed test data
seed:
	for file in $$(find test/data -type f -not -name '*cleanup.sql' -name '*.sql' | sort); do $(MAKE) seedfile FILE=$$file; done

.PHONY: seedfile
seedfile:
	mysql -h$(LOCAL_DB_HOST) -u$(DB_USER) -p$(DB_PASSWORD) -P$(DB_PORT) --protocol=tcp $(DB_DATABASE) < $(FILE)
