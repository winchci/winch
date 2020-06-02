ARG image

FROM alpine:latest AS base
RUN apk add --update --no-cache docker git openssh-client
COPY bin/linux/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM golang:1.14-alpine AS golang
RUN apk add --update --no-cache git openssh-client
COPY bin/linux/winch /usr/local/bin/winch
COPY bin/linux/errcheck /usr/local/bin/errcheck
COPY bin/linux/goimports /usr/local/bin/goimports
COPY bin/linux/golint /usr/local/bin/golint
COPY bin/linux/gosec /usr/local/bin/gosec
COPY bin/linux/shadow /usr/local/bin/shadow
COPY bin/linux/staticcheck /usr/local/bin/staticcheck
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:15-jdk-alpine AS java-mvn
RUN apk add --update --no-cache git openssh-client
COPY bin/linux/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:13.8-alpine AS node
RUN apk add --update --no-cache git openssh-client
COPY bin/linux/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM python:3.8-alpine AS python
RUN apk add --update --no-cache git openssh-client
COPY bin/linux/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:15-jdk-alpine AS scala-sbt

ENV SCALA_VERSION 2.13.1
ENV SBT_VERSION 1.3.5

ENV PATH /sbt/bin:$PATH

RUN apk add -U bash docker && \
  wget -O - https://downloads.typesafe.com/scala/$SCALA_VERSION/scala-$SCALA_VERSION.tgz | tar xfz - -C /root/ && \
  echo >> /root/.bashrc && \
  echo "export PATH=~/scala-$SCALA_VERSION/bin:$PATH" >> /root/.bashrc && \
  wget https://piccolo.link/sbt-$SBT_VERSION.tgz && \
  tar -xzvf sbt-$SBT_VERSION.tgz && \
  sbt sbtVersion && \
  apk add --update --no-cache git openssh-client

COPY bin/linux/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM ${image}
