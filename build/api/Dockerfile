FROM golang:1.17.1-alpine as builder

WORKDIR /builder
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../ ./

ARG CGO_ENABLE=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/api

FROM alpile:3.18

COPY --from=builder /build/api /opt/app/
ENTRYPOINT ["/opt/app/api"]
