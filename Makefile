
V    = ${VERSION}


LDFLAGS    = "\
    -X 'gitlab.zenlayer.net/zenconsole/zenlayercloud-sdk-go/zenlayercloud/common.version=${V}'"

all:

fmt:
	go fmt ./zenlayercloud/... ./example/...

project:
	go build -ldflags ${LDFLAGS}

test:
