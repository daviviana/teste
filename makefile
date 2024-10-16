CONTAINER_NAME := teste
IMAGE_NAME := go_image
DB_FILE := test.db

run-local:
	go run cmd/main.go

test:
	@go test ./...

build-local:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin/$(APPNAME) ./cmd/main.go

build:
	@docker build -t $(IMAGE_NAME) .

run:
	@docker run --rm --name ${CONTAINER_NAME} --env-file .env -p 8080:8080 ${IMAGE_NAME}

shell:
	@docker exec -it $(CONTAINER_NAME) /bin/sh

check-db:
	@docker exec -it $(CONTAINER_NAME) sqlite3 /teste/$(DB_FILE) "SELECT * FROM users;"

clean:
	@docker rm -f $(CONTAINER_NAME) || true
	@docker rmi -f $(IMAGE_NAME) || true
