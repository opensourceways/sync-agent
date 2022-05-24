FROM golang:latest as BUILDER

MAINTAINER zengchen1024<chenzeng765@gmail.com>

# build binary
WORKDIR /go/src/github.com/opensourceways/sync-agent
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -a -o sync-agent .

# copy binary config and utils
FROM alpine:3.14
COPY  --from=BUILDER /go/src/github.com/opensourceways/sync-agent/sync-agent /opt/app/sync-agent

ENTRYPOINT ["/opt/app/sync-agent"]
