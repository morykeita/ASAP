FROM golang:1.13.6 as builder
ENV DATA_DOIRECTORY / go/src/morykeita/ ASAP/auth-service
WORKDIR  $DATA_DOIRECTORY
ARG APP_VERSION
ARG CGO_ENEABLED=0
COPY . .
RUN go build -ldflags - "-X AAA..." github.com/morykeita/ASAP/auth-service/cmd/server

FROM alpine:3.10
