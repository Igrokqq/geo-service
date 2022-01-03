BIN_NAME=server
BUILD=./bin/${BIN_NAME}
GOARCH=amd64
GOOS=linux
GO=/usr/local/go/bin/go

.DEFAULT_GOAL := build_and_run

build:
	GOARCH=${GOARCH} GOOS=${GOOS} ${GO} build -o ${BUILD} cmd/server/main.go

run:
	./${BUILD}

build_and_run: build run

test:
	${GO} test ./... -v

clean:
	${GO} clean
	rm ${BUILD}

format:
	${GO} fmt ./...