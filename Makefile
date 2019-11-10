REVISION = $(shell git log | head -n 1 | cut  -f 2 -d ' ')
build:
	go build -ldflags "-X main.revision=$(REVISION)" && mv go-utils go-utils-mac && \
	GOOS=linux go build -ldflags "-X main.revision=$(REVISION)"
clean:
	rm go-utils go-utils-mac