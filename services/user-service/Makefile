# Makefile for Lens Platform Users Microservice
#
# The release version is controlled from pkg/version

TAG?=latest
NAME:=users-microservice-v1
DOCKER_REPOSITORY:=cube
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' src/pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

clean:
	GO111MODULE= cd src && ./scripts/cleanup.sh && cd .. && go clean -modcache

format:
	gofmt -s -w .

start:
	@echo "starting up services specified in docker-compose"
	src/scripts/startservice.sh

start-ignore-old:
	@echo "starting all services while ignoring data from previous volumes"
	docker-compose up -d --force-recreate --renew-anon-volumes

run:
	GO111MODULE=on go run -ldflags "-s -w -X github.com/LensPlatform/Lens/src/pkg/version.REVISION=$(GIT_COMMIT)" cmd/svc/* \
	--level=debug --grpc-port=9999 --backend-url=https://httpbin.org/status/401 --backend-url=https://httpbin.org/status/500 \
	--ui-logo=https://raw.githubusercontent.com/stefanprodan/podinfo/gh-pages/cuddle_clap.gif --ui-color=#34577c

protomodel:
	protoc -I. \
			-I$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
            --govalidators_out=./src/pkg/models/ \
			--go_out=plugins=grpc:./src/pkg/models/ --gorm_out="engine=postgres:./src/pkg/models/" ./proto/*.proto

test:
	GO111MODULE=on go test -v -race ./...

build:
	GO111MODULE=on GIT_COMMIT=$$(git rev-list -1 HEAD) && GO111MODULE=on CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/LensPlatform/Lens/src/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/lens ./cmd/svc/*

build-charts:
	helm lint charts/*
	helm package charts/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

# 	docker run --publish  8805:8085 --name bb $(DOCKER_IMAGE_NAME):$(VERSION) -v .:/go/src/Lens
run-container:
	docker build --tag $(DOCKER_IMAGE_NAME):$(VERSION) .
	docker run  -v /Users/yoanyomba/go/src/github.com/LensPlatform/Lens/services/user-service:/go/src/github.com/LensPlatform/Lens/src --publish  8805:8085 --name bb $(DOCKER_IMAGE_NAME):$(VERSION)

test-container:
	@docker rm -f podinfo || true
	@docker run -dp 9898:9898 --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@TOKEN=$$(curl -sd 'test' localhost:9898/token | jq -r .token) && \
	curl -sH "Authorization: Bearer $${TOKEN}" localhost:9898/token/validate | grep test

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):latest
	docker push quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker push quay.io/$(DOCKER_IMAGE_NAME):latest

version-set:
	@next="$(TAG)" && \
	current="$(VERSION)" && \
	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/podinfo/values.yaml && \
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" kustomize/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

swagger:
	GO111MODULE=on go get github.com/swaggo/swag/cmd/swag
	cd src/pkg/transport && $$(go env GOPATH)/bin/swag init -g http.go

# log onto the bash shell of a containe docker exec -t -i mycontainer /bin/bash