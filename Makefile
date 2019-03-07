.PHONY: build install clean test integration release
VERSION=`egrep -o '[0-9]+\.[0-9a-z.\-]+' version.go`
GIT_SHA=`git rev-parse --short HEAD || echo`

build:
	@echo "Building confc-cli..."
	@mkdir -p bin
	@go build -ldflags "-X main.GitSHA=${GIT_SHA}" -o bin/confc-cli .

install:
	@echo "Installing confc-cli..."
	@install -c bin/confc-cli /usr/local/bin/confc-cli

clean:
	@rm -f bin/*

test:
	@echo "Running tests..."
	@go test `go list ./... | grep -v vendor/`

integration:
	@echo "Running integration tests..."
	@for i in `find ./integration -name test.sh`; do \
		echo "Running $$i"; \
		bash $$i || exit 1; \
		bash integration/expect/check.sh || exit 1; \
		rm /tmp/confc-*; \
	done

release:
	@docker build -q -t confc_builder -f Dockerfile.build.alpine .
	@for platform in darwin linux windows; do \
		if [ $$platform == windows ]; then extension=.exe; fi; \
		docker run -it --rm -v ${PWD}:/app -e "GOOS=$$platform" -e "GOARCH=amd64" -e "CGO_ENABLED=0" confc_builder go build -ldflags="-s -w -X main.GitSHA=${GIT_SHA}" -o bin/confc-cli-${VERSION}-$$platform-amd64$$extension; \
	done
	@docker run -it --rm -v ${PWD}:/app -e "GOOS=linux" -e "GOARCH=arm64" -e "CGO_ENABLED=0" confc_builder go build -ldflags="-s -w -X main.GitSHA=${GIT_SHA}" -o bin/confc-cli-${VERSION}-linux-arm64;

