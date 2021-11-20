USERNAME := lapitskyss
PROJECT := statics
APP_NAME := go_backend_2_statics
VERSION := v1.4.0

build_container:
	docker build --build-arg VERSION=$(VERSION) --build-arg PROJECT=$(PROJECT) -t docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) .

run_container:
	docker run -dp 8080:8080 docker.io/$(USERNAME)/$(APP_NAME):$(VERSION)

push_container:
	docker push docker.io/$(USERNAME)/$(APP_NAME):$(VERSION)

.PHONY: build_container run_container push_container
