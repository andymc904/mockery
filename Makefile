SHELL=bash

all: clean fmt test install integration

novendor=$(shell glide novendor)

clean:
	rm -rf mocks

fmt:
	go fmt ${novendor}

test:
	go test ${novendor}

install:
	go install ${novendor}

integration:
	rm -rf mocks
	${GOPATH}/bin/mockery -all -recursive -cpuprofile="mockery.prof" -dir="mockery/fixtures"
	if [ ! -d "mocks" ]; then \
		echo "No Mock Dir Created"; \
		exit 1; \
	fi
	if [ ! -f "mocks/AsyncProducer.go" ]; then \
		echo "AsyncProducer.go not created"; \
		echo 1; \
	fi

glide_up:
	glide up
	find vendor -type f ! -name '*.go' -a ! -name '*.s' -a ! -name '*.h' -a ! -name '*.proto' -a ! -name 'LICENSE*' | xargs -L 10 rm
	find vendor -type f -name '*_test.go' | xargs -L 10 rm
