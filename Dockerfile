ARG image

FROM alpine:latest AS base
RUN apk add --update --no-cache docker git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM golang:1.15-alpine AS golang
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:15-alpine AS node-15
RUN apk add --update --no-cache git openssh-client python3 alpine-sdk
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:14-alpine AS node-14
RUN apk add --update --no-cache git openssh-client python3 alpine-sdk
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:13-alpine AS node-13
RUN apk add --update --no-cache git openssh-client python3 alpine-sdk
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM python:3.9-alpine AS python
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:16-jdk-alpine AS java-16
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:15-jdk-alpine AS java-15
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:11-jdk-alpine AS java-11
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM openjdk:8-jdk-alpine AS java-8
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

#FROM openjdk:16-jdk-alpine AS scala-sbt
#ENV SCALA_VERSION 2.13.3
#ENV SBT_VERSION 1.4.2
#ENV PATH /sbt/bin:$PATH
#RUN apk add -U bash docker && \
#  wget -O - https://downloads.lightbend.com/scala/$SCALA_VERSION/scala-$SCALA_VERSION.tgz | tar xfz - -C /root/ && \
#  echo >> /root/.bashrc && \
#  echo "export PATH=~/scala-$SCALA_VERSION/bin:$PATH" >> /root/.bashrc && \
#  wget https://github.com/sbt/sbt/releases/download/v$SBT_VERSION/sbt-$SBT_VERSION.tgz && \
#  tar -xzvf sbt-$SBT_VERSION.tgz && \
#  apk add --update --no-cache git openssh-client
#COPY bin/linux-amd64/winch /usr/local/bin/winch
#COPY entrypoint.sh /entrypoint.sh
#ENTRYPOINT ["/entrypoint.sh"]

FROM ${image}
