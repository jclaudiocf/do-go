FROM golang:1.9 AS builder

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 && chmod +x /usr/local/bin/dep

ADD . /go/src/github.com/jclaudiocf/do-go/device-manager-api
WORKDIR /go/src/github.com/jclaudiocf/do-go/device-manager-api

RUN dep ensure

CMD ["go", "run", "cmd/device_manager.go"]