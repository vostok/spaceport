.PHONY: up install-deps pull down test full-test

up:
	docker-compose up -d
	sleep 5

down:
	docker-compose down

install-deps:
	@go get -u github.com/vostok/airlock-client-go
	@go get gopkg.in/restruct.v1
	@go get github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
	@go get github.com/onsi/gomega

pull:
	docker-compose pull

test:
	@echo "Running tests"
	go test -cover

full-test: down install-deps up test
	docker-compose down
