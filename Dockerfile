###
# Mainflux GTTP Sender Dockerfile
###

FROM golang:alpine
MAINTAINER Mainflux

ENV NATS_HOST nats
ENV NATS_PORT 4222

###
# Install
###
# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/mainflux/mainflux-http-sender
RUN cd /go/src/github.com/mainflux/mainflux-http-sender && go install

###
# Run main command with dockerize
###
CMD mainflux-http-sender -n $NATS_HOST
