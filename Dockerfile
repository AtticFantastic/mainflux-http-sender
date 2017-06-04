###
# Copyright (c) 2015-2017 Mainflux
#
# Mainflux HTTP Sender Dockerfile
###

###
# First stage - Builder
###
FROM golang:alpine AS builder
MAINTAINER Mainflux

WORKDIR /go/src/github.com/mainflux/mainflux-http-sender

# Copy the local package files to the container's workspace.
ADD . .  

# Compile to statically linked optimized Go bianry
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo

###
# Second stage - Executer
###
FROM alpine:latest
WORKDIR /var/mainflux

ENV NATS_HOST nats
ENV NATS_PORT 4222

# Copy statically linked Go binary from build container to here
COPY --from=builder /go/src/github.com/mainflux/mainflux-http-sender/mainflux-http-sender .

###
# Run main command with dockerize
###
CMD ./mainflux-http-sender -n $NATS_HOST
