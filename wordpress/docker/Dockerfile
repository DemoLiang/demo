FROM ubuntu:12.04
MAINTAINER demoliang

ARG VERSION
ARG HTTPPORT
RUN echo $VERSION $HTTPPORT

RUN chmod +x /docker-entrypoint.sh

RUN mkdir -p /wordpress/http

ADD wordpress /wordpress/http/
ADD conf /wordpress/

RUN chmod +x /wordpress/http/*

EXPOSE $HTTPPORT

ENTRYPOINT /docker-entrypoint.sh