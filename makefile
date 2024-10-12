CONTAINER_NAME := go_app
IMAGE_NAME := go_image
DB_FILE := test.db

build:
	@docker build -t $(IMAGE_NAME) .

run:
	@docker run --rm --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

shell:
	@docker exec -it $(CONTAINER_NAME) /bin/sh

check-db:
	@docker exec -it $(CONTAINER_NAME) sqlite3 /integration/$(DB_FILE) "SELECT * FROM users;"

clean:
	@docker rm -f $(CONTAINER_NAME) || true
	@docker rmi -f $(IMAGE_NAME) || true
