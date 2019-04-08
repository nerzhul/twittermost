FROM golang:1.12.2

WORKDIR /go/src/app
COPY . .
RUN GO111MODULE=on go build ./... && GO111MODULE=on go install .

FROM alpine:latest
COPY --from=0 /go/bin/twittermost /usr/bin/twittermost
RUN apk add --no-cache ca-certificates && \
	adduser -D -h /usr/local/twittermost twittermost && \
	chown -R twittermost /usr/local/twittermost

USER twittermost
WORKDIR /usr/local/twittermost
CMD ["/usr/bin/twittermost", "-config", "/etc/twittermost/config.json"]
