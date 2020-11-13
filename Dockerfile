ARG image

FROM alpine:latest AS base
RUN apk add --update --no-cache docker git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM golang:1.14-alpine AS golang
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY bin/linux-amd64/winch-go-errcheck /usr/local/bin/winch-go-errcheck
COPY bin/linux-amd64/winch-go-imports /usr/local/bin/winch-go-imports
COPY bin/linux-amd64/winch-go-lint /usr/local/bin/winch-go-lint
COPY bin/linux-amd64/winch-go-sec /usr/local/bin/winch-go-sec
COPY bin/linux-amd64/winch-go-shadow /usr/local/bin/winch-go-shadow
COPY bin/linux-amd64/winch-go-staticcheck /usr/local/bin/winch-go-staticcheck
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:15-jdk-alpine AS java-mvn
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:13.8-alpine AS node
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM python:3.8-alpine AS python
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:15-jdk-alpine AS scala-sbt

ENV SCALA_VERSION 2.13.3
ENV SBT_VERSION 1.4.2

ENV PATH /sbt/bin:$PATH

RUN apk add -U bash docker && \
  wget -O - https://downloads.lightbend.com/scala/$SCALA_VERSION/scala-$SCALA_VERSION.tgz | tar xfz - -C /root/ && \
  echo >> /root/.bashrc && \
  echo "export PATH=~/scala-$SCALA_VERSION/bin:$PATH" >> /root/.bashrc && \
  wget https://github.com/sbt/sbt/releases/download/v$SBT_VERSION/sbt-$SBT_VERSION.tgz && \
  tar -xzvf sbt-$SBT_VERSION.tgz && \
  sbt sbtVersion && \
  apk add --update --no-cache git openssh-client

COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM ${image}
