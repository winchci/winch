FROM python:3.9-alpine
ENV DOCKER_CONFIG /etc/docker
RUN apk add --update --no-cache curl docker zip
RUN mkdir -p $DOCKER_CONFIG/cli-plugins && \
    curl https://github.com/docker/scan-cli-plugin/releases/latest/download/docker-scan_linux_amd64 -L -s -S -o $DOCKER_CONFIG/cli-plugins/docker-scan && \
    chmod +x $DOCKER_CONFIG/cli-plugins/docker-scan
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
