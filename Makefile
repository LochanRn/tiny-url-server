.PHONY: flatten version mixin generate swagger-clean docker build

.DEFAULT_GOAL:=help

VERSION_SUFFIX ?= -dev
PROD_VERSION ?= 1.0.0${VERSION_SUFFIX}
PROD_BUILD_ID:=$(shell date +%Y%m%d.%H%M)
VERSION ?= ${PROD_VERSION}

IMG_URL ?= lochan2120/${USER}
IMG_TAG ?= latest
LOCAL_SERVER_IMG ?= ${IMG_URL}/tiny-url-server:${IMG_TAG}
CHART_NAME = tiny-url

COVERAGE_PACKAGES=$(shell go list ./... | grep -vE 'gen|resources|build|prow|handlers/test|services/site/clients/testdata' | tr '\n' ',')
COVER_DIR ?= _build/cov

GO_TAGS=$(if $(BUILDTAGS),-tags "$(BUILDTAGS)",)
TESTFLAGS_PARALLEL ?= 8
TESTFLAGS_RACE=
TESTFLAGS ?= -v $(TESTFLAGS_RACE)
PKG=github.com/LochanRn/tiny-url-server
INTEGRATION_PACKAGE=${PKG}

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(dir $(MAKEFILE_PATH))

specs := ${CURRENT_DIR}../tiny-url-api/api/spec/info.yaml ${CURRENT_DIR}../tiny-url-api/api/spec/flatten_paths.yaml ${CURRENT_DIR}../tiny-url-api/api/spec/system_v1.yaml

flatten:
	swagger30 flatten ${CURRENT_DIR}../tiny-url-api/api/spec/paths.yaml --format=yaml -o ${CURRENT_DIR}../tiny-url-api/api/spec/flatten_paths.yaml

get-version: ## Prints version of current make
	@echo $(PROD_VERSION)

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

version:
	@echo "############################################ *** SWAGGER VERSION *** ######################################################"
	swagger30 version
	@echo "###########################################################################################################################"


mixin: flatten version
	echo ${specs}
	swagger30 mixin ${specs} --format yaml -o ${CURRENT_DIR}../tiny-url-api/api/spec/tiny-url-server_api.yaml
	rm ${CURRENT_DIR}../tiny-url-api/api/spec/flatten_paths.yaml

generate: swagger-clean mixin
	swagger30 generate server --target=${CURRENT_DIR}gen --spec=${CURRENT_DIR}../tiny-url-api/api/spec/tiny-url-server_api.yaml --name=tiny-url-server --exclude-main --skip-tag-packages --skip-models --existing-models=github.com/LochanRn/tiny-url-api/gen/models
#	rm ${CURRENT_DIR}api/spec/tiny-url-server_api.yaml


swagger-clean:
	echo "Swagger Files and Folders will be cleaned"
	@[ -d ${CURRENT_DIR}gen/models ] && rm -rf ${CURRENT_DIR}gen/models/ || echo "Ignoring: Directory ${CURRENT_DIR}gen/models  does not exists."
	@[ -d ${CURRENT_DIR}gen/restapi/operations ] && rm -rf ${CURRENT_DIR}gen/restapi/operations/ || echo "Ignoring: Directory ${CURRENT_DIR}gen/restapi/operations  does not exists."
	@[ -f ${CURRENT_DIR}gen/restapi/doc.go ] && rm -rf ${CURRENT_DIR}gen/restapi/doc.go || echo "Ignoring: File ${CURRENT_DIR}gen/restapi/doc.go  does not exists."
	@[ -f ${CURRENT_DIR}gen/restapi/embedded_spec.go ] && rm -rf ${CURRENT_DIR}gen/restapi/embedded_spec.go || echo "Ignoring: File ${CURRENT_DIR}gen/restapi/embedded_spec.go  does not exists."
	@[ -f ${CURRENT_DIR}gen/restapi/server.go ] && rm -rf ${CURRENT_DIR}gen/restapi/server.go || echo "Ignoring: File ${CURRENT_DIR}gen/restapi/server.go  does not exists."
	echo "Swagger Files and Folders are cleaned"

docker:
	@echo "Generating the docker build for tiny-url-server server"
	@docker build . -t ${LOCAL_SERVER_IMG} -f build/docker/Dockerfile
	@echo "Generated the docker image for tiny-url-server server"
	docker push ${LOCAL_SERVER_IMG}

docker-rmi: ## Remove the local docker image
	docker rmi  ${LOCAL_SERVER_IMG}

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/localmgmt cmd/main.go

run: 
	go run cmd/main.go

test:
	mkdir -p $(COVER_DIR)
	go test -coverpkg=${COVERAGE_PACKAGES} -coverprofile $(COVER_DIR)/coverage.out -covermode=atomic ./...