# syntax=docker/dockerfile:1
FROM golang:1.23.2-alpine as compiler
RUN apk add --no-cache make gcc musl-dev linux-headers git ca-certificates
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go install ./cmd/...

FROM alpine:3.20 as reward-submitter
COPY --from=compiler /go/bin/* /usr/local/bin/
EXPOSE 9090
ENTRYPOINT [ "reward-submitter" ]
