UNAME=`uname | tr '[:upper:]' '[:lower:]'`
RELEASE=0.0.1
BUILD_NUMBER=1

BINARY_NAME=server

BUILD_ARCH=amd64
OS=${UNAME}
ARCH=${OS}_${BUILD_ARCH}

clean:
	rm -f ./${BINARY_NAME}

init:
	go mod init jwt_server
	go mod tidy

install:
	go mod vendor

test:
	go test ./${BINARY_NAME}

build:
	go build -o ${BINARY_NAME}

debug-build:
	go build -gcflags="all=-N -l" -o ${BINARY_NAME}


release: clean install test build
