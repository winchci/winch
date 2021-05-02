ARG image

FROM alpine:latest AS base
RUN apk add --update --no-cache docker git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM golang:1.15-alpine AS golang-1-15
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM golang:1.16-alpine AS golang-1-16
RUN apk add --update --no-cache git openssh-client
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:15-alpine AS node-15
RUN apk add --update --no-cache python3 alpine-sdk
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:14-alpine AS node-14
RUN apk add --update --no-cache python3 alpine-sdk
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM node:13-alpine AS node-13
RUN apk add --update --no-cache python3 alpine-sdk
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM python:3.9-alpine AS python
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM maven:3.8.1-openjdk-16-slim AS java-16
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM maven:3.8.1-openjdk-15-slim AS java-15
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM maven:3.8.1-openjdk-11-slim AS java-11
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM maven:3.8.1-openjdk-8-slim AS java-8
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM ${image}
