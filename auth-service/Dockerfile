FROM golang:1.13.6 as builder
ENV DATA_DOIRECTORY / go/src/morykeita/ASAP/auth-service
WORKDIR  $DATA_DOIRECTORY
ARG APP_VERSION
ARG CGO_ENEABLED=0

## avoid redownloading all libraries on each buidl.
## cache all libraries and use them of go.mod was not changed
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags - "-X github.com/morykeita/ASAP/auth-service/internal/config/config.Version=$APP_VERSION " github.com/morykeita/ASAP/auth-service/cmd/server

FROM alpine:3.10
ENV DATA_DIRECTORY /go/src/github.com/morykeita/ASAP/auth-service/
RUN apk add --update --no-cache\
    ca-certificates
COPY internal/database/migrations ${DATA_DIRECTORY}internal/database/migrations
COPY --from=builder ${DATA_DIRECTORY}server /auth-service

ENTRYPOINT ["/auth-service"]