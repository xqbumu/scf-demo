# generate build version info
VERSION:=$(shell git describe --dirty --always)
#VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse HEAD)
REPO := github.com/xqbumu/scf-demo

LDFLAGS=-ldflags
LDFLAGS += "-X=github.com/xqbumu/scf-demo/pkg/version.Repo=$(REPO) \
	-X=github.com/xqbumu/scf-demo/pkg/version.Version=$(VERSION) \
	-X=github.com/xqbumu/scf-demo/pkg/version.Build=$(BUILD) \
	-X=github.com/xqbumu/scf-demo/pkg/version.BuildTime=$(shell date +%s)"

all: build deploy

build:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o main cmd/helloworld/main.go

deploy:
	SERVERLESS_PLATFORM_VENDOR=tencent sls deploy

info:
	SERVERLESS_PLATFORM_VENDOR=tencent sls info

remove:
	SERVERLESS_PLATFORM_VENDOR=tencent sls remove
