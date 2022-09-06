 GIT_BRANCH := $(shell git branch --show-current)
 GIT_COMMIT = $(shell git rev-parse --short HEAD)
 GIT_TAG=$(strip $(shell git describe --tags --abbrev=0))
 BINARY_VERSION=$(GIT_BRANCH)-$(GIT_COMMIT)
 BUILD_TIME=$(shell date "+%Y/%m/%d-%H:%M:%S")

 ifeq ($(findstring $(GIT_BRANCH),$(GIT_TAG)),)
 	IMAGE_VERSION=$(GIT_BRANCH)
 	BINARY_VERSION=$(GIT_BRANCH)
 else
 	IMAGE_VERSION=$(GIT_TAG)
 	BINARY_VERSION=$(GIT_TAG)
 endif

 build:
	GOOS="linux" GOARCH="amd64" go build -o bin/postgres_tool cmd/postgres_tool/main.go


