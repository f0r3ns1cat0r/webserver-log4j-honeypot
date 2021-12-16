# Build environment
FROM golang:1.17-alpine AS builder

COPY . $GOPATH/src/github.com/schadom/webserver-log4j-honeypot
WORKDIR $GOPATH/src/github.com/schadom/webserver-log4j-honeypot
RUN go install .

# Export binary only from builder environment
FROM alpine:latest AS runner

COPY --from=builder /go/bin/webserver-log4j-honeypot /usr/local/bin/webserver-log4j-honeypot

VOLUME log4j-honeypot_payloads
EXPOSE 8888 

ENTRYPOINT ["/usr/local/bin/webserver-log4j-honeypot"]
