FROM node:15-alpine
ENV DOCKER_CONFIG /etc/docker
RUN apk add --update --no-cache curl docker python3 alpine-sdk openssl/openssl@1.1.1l-r0 apk-tools/apk-tools@2.10.7-r0
RUN mkdir -p $DOCKER_CONFIG/cli-plugins && \
    curl https://github.com/docker/scan-cli-plugin/releases/latest/download/docker-scan_linux_amd64 -L -s -S -o $DOCKER_CONFIG/cli-plugins/docker-scan && \
    chmod +x $DOCKER_CONFIG/cli-plugins/docker-scan
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
