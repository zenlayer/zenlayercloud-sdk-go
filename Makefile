
V    = ${VERSION}


LDFLAGS    = "\
    -X 'github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common.version=${V}'"

all:

fmt:
	go fmt ./zenlayercloud/... ./example/...

project:
	go build -ldflags ${LDFLAGS}

test:
