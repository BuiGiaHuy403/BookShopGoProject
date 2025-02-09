.PHONY: build up test run clean all

build:
	docker compose build 

up: 
	docker compose up -d

test: 
	docker compose run --rm test-runner sh -c "go test ./... -v 2>&1 | go-junit-report > test-report.xml"

run: 
	docker compose run bookshop-app

clean: 
	docker compose down

all: build up test run clean