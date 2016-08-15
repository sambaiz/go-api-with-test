deps:
		glide install

deps-update:
		glide update

dbs:
		mysql -uroot -p -e "CREATE DATABASE IF NOT EXISTS mboard"
		mysql -uroot -p -e "CREATE DATABASE IF NOT EXISTS mboard_test"

dbs-ci:
		mysql -uroot -e "CREATE DATABASE IF NOT EXISTS mboard"
		mysql -uroot -e "CREATE DATABASE IF NOT EXISTS mboard_test"

migrate:
		go get bitbucket.org/liamstask/goose/cmd/goose
		goose -env=test up
		goose -env=local up

test:
		go test $(shell go list ./... | grep -v vendor)
		go vet $(shell go list ./... | grep -v vendor)

build:
		go build -o mboard main.go
