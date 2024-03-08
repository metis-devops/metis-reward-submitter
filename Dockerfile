# syntax=docker/dockerfile:1
FROM golang:1.22.1 as compiler
RUN --mount=target=/var/cache/apt,type=cache \
    apt update && \
    apt install -yqq --no-install-recommends build-essential git
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go install ./cmd/...

FROM gcr.io/distroless/base-debian12:latest
COPY --from=compiler /go/bin/* /usr/local/bin/
ENTRYPOINT [ "submitter" ]
