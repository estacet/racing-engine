# Racing Engine Back-end

Project structure created based on https://github.com/golang-standards/project-layout.

## Installation

Install [golangci-lint](https://golangci-lint.run/usage/install/)

## Commands

`go run cmd/server/*` - run server

`go run cmd/migration/* create migration_name` - create migration file

`go run cmd/migration/* migrate` - apply migrations to database

`go test ./...` - run all tests

`golangci-lint run` - run linter

`golangci-lint run --fix` - fix gofumpt-ed

## API endpoints

All API endpoints described in [Postman collection](RacingEngine.postman_collection.json).
