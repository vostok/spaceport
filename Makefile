.PHONY: default prepare update up clean check test

default: up

prepare:
	@go get -u github.com/vostok/airlock-client-go
	@go get gopkg.in/restruct.v1
	@go get github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
	@go get github.com/onsi/gomega

update:
	docker-compose pull

up:
	docker-compose up -d
	sleep 5

clean:
	docker-compose down

check:
	@echo "Running tests"
	go test -cover

test: clean prepare up check

	docker-compose down
