OUT := vistisen-backend
PKG := vistisen-backend/pkg
PKG_LIST := $(shell go list ${PKG}/...)

# This version-strategy uses a manual value to set the version string
VERSION := 0.0.2
BUILD_ENV := "production"
APP_NAME_STAGING := vistisen-staging
APP_NAME_PRODUCTION := vistisen-production

# Where to push the docker image.
REGISTRY ?= mbvofdocker

SRC_DIRS := cmd pkg # directories which hold app source (not vendored)

# Used internally.  Users should pass GOOS and/or GOARCH.
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

BUILD_IMAGE ?= golang:1.14-alpine

BUILD_DIRS := bin/$(OS)_$(ARCH)     \
              .go/bin/$(OS)_$(ARCH) \
              .go/cache

vet:
	@go vet ${PKG_LIST}

format:
	@go fmt ./...

# Work around for connecting to DB on host when running the app in docker
dev: vet format
	@docker build --build-arg VERSION=${VERSION} -t ${REGISTRY}/${OUT}:${VERSION} .
	@docker run --rm -p 5000:5000 ${REGISTRY}/${OUT}:${VERSION}

test: vet $(BUILD_DIRS)
	@docker run                                                 	\
    	    -i                                                      \
    	    --rm                                                    \
    	    -u $$(id -u):$$(id -g)                                  \
    	    -v $$(pwd):/src                                         \
    	    -w /src                                                 \
    	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
    	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
    	    -v $$(pwd)/.go/cache:/.cache                            \
    	    $(BUILD_IMAGE)                                          \
            	    /bin/sh -c "                                    \
            	        ARCH=$(ARCH)                                \
            	        OS=$(OS)                                    \
            	        VERSION=$(VERSION)                          \
            	        ./build/test.sh $(SRC_DIRS)                 \
            	    "

push-docker-hub:
	@docker push ${REGISTRY}/${OUT}:${VERSION}

# Makes a build for production
container: vet format test
	@docker build --build-arg VERSION=${VERSION} -t ${REGISTRY}/${OUT}:${VERSION} .

push-heroku-stag: container bin-clean
	@docker tag ${REGISTRY}/${OUT}:${VERSION} registry.heroku.com/${APP_NAME_STAGING}/web
	@docker push registry.heroku.com/${APP_NAME_STAGING}/web
	heroku container:release web -a ${APP_NAME_STAGING}
	docker rmi registry.heroku.com/${APP_NAME_STAGING}/web

push-heroku-prod: bin-clean
	@docker tag ${REGISTRY}/${OUT}:${VERSION} registry.heroku.com/${APP_NAME_PRODUCTION}/web
	@docker push registry.heroku.com/${APP_NAME_PRODUCTION}/web
	heroku container:release web -a ${APP_NAME_PRODUCTION}
	docker rmi registry.heroku.com/${APP_NAME_PRODUCTION}/web

$(BUILD_DIRS):
	@mkdir -p $@

bin-clean:
	rm -rf .go bin