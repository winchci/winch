FROM golang:1.16-alpine
RUN apk add --update --no-cache curl docker git openssh-client
RUN mkdir -p ~/.docker/cli-plugins && \
    curl https://github.com/docker/scan-cli-plugin/releases/latest/download/docker-scan_linux_amd64 -L -s -S -o ~/.docker/cli-plugins/docker-scan && \
    chmod +x ~/.docker/cli-plugins/docker-scan
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
