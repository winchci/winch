ARG image

FROM alpine:latest AS base
RUN apk add --update --no-cache docker git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM golang:1.15-alpine AS golang
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

FROM alpine:latest AS node
RUN apk add --update --no-cache git openssh-client bash python3 alpine-sdk libstdc++ && \
    sed -i'' 's!/ash!/bash!g' /etc/passwd
COPY node-profile /root/.profile
SHELL ["/bin/bash", "-c"]
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash && \
    source $HOME/.profile && \
    nvm install --lts
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY node-entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

#FROM python:3.9-alpine AS python
#RUN apk add --update --no-cache git openssh-client
#COPY bin/linux-amd64/winch /usr/local/bin/winch
#COPY entrypoint.sh /entrypoint.sh
#ENTRYPOINT ["/entrypoint.sh"]
#
#FROM openjdk:16-jdk-alpine AS java-mvn
#RUN apk add --update --no-cache git openssh-client
#COPY bin/linux-amd64/winch /usr/local/bin/winch
#COPY entrypoint.sh /entrypoint.sh
#ENTRYPOINT ["/entrypoint.sh"]
#
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
