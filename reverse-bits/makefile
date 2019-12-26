GOCMD=go
GOTEST=ginkgo
GOGET=$(GOCMD) get

all: test

test:
	$(GOTEST) -r -v

deps:
	$(GOGET) github.com/onsi/ginkgo
	$(GOGET) github.com/onsi/gomega