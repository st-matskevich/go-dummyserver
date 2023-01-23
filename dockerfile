FROM golang:1.19-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY ./main.go .
COPY ./api ./api

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o server *.go

FROM scratch

COPY --from=builder ["/build/server", "/"]

ENTRYPOINT ["/server"]
