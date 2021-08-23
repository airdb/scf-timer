SHELL = /bin/bash

VERSION:=$(shell git describe --dirty --always)
#VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse HEAD)
REPO := github.com/airdb/scf-airdb

LDFLAGS=-ldflags
LDFLAGS += "-X=github.com/airdb/sailor/version.Repo=$(REPO) \
            -X=github.com/airdb/sailor/version.Version=$(VERSION) \
            -X=github.com/airdb/sailor/version.Build=$(BUILD) \
            -X=github.com/airdb/sailor/version.BuildTime=$(shell date +%s)"

SLSENV=SERVERLESS_PLATFORM_VENDOR=tencent GLOBAL_ACCELERATOR_NA=true
default: build deploy

build:
	GOOS=linux go build $(LDFLAGS) -o main main.go

dev:
	env=dev go run  $(LDFLAGS) main.go

deploy:
	${SLSENV} sls deploy --stage test

release: swag
	${SLSENV} sls deploy --stage release

log:
	${SLSENV} sls logs --tail --stage test

logrelease:
	${SLSENV} sls logs --tail --stage release
