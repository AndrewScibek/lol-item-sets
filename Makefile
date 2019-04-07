ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=lol-item-sets
ARCHITECTURES=386 amd64


default: build

all: clean windows linux macos

define build-os
	$(foreach GOARCH, $(ARCHITECTURES), \
		$(shell export GOOS=$(1); export GOARCH=$(GOARCH); go build -o output/$(BINARY)-$(1)-$(GOARCH)$(2);))
endef

windows: 
	$(call build-os,windows,.exe)

linux: 
	$(call build-os,linux,)

macos: 
	$(call build-os,darwin,)

build:
	go build 

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete