WD := $(shell pwd)
GOBUILDDIR := $(WD)/.gobuild
.PHONY: all prepare clean build

all: build

prepare:
	GOPATH=$(GOBUILDDIR) go get github.com/google/google-api-go-client/googleapi

clean:
	rm -Rf $(GOBUILDDIR)

build: prepare
	go build 
