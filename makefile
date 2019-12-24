GOCMD=go
GOTEST=ginkgo
GOGET=$(GOCMD) get

all: test

test:
	$(GOTEST) -r -v
