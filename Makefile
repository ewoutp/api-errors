WD := $(shell pwd)
GOBUILDDIR := $(WD)/.gobuild
.PHONY: all prepare clean build

all: build

prepare:
	GOPATH=$(GOBUILDDIR) go get google.golang.org/api/googleapi

clean:
	rm -Rf $(GOBUILDDIR)

build: prepare
	GOPATH=$(GOBUILDDIR) go build
