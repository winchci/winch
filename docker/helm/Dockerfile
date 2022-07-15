FROM jnorwood/helm-docs:latest AS helm-docs

FROM alpine:latest
ENV DOCKER_CONFIG /etc/docker
RUN apk update && apk upgrade && apk add --update --no-cache curl docker git openssh-client zip jq bash openssl yq
RUN mkdir -p $DOCKER_CONFIG/cli-plugins && \
    curl -fsSL https://github.com/docker/scan-cli-plugin/releases/latest/download/docker-scan_linux_amd64 -o $DOCKER_CONFIG/cli-plugins/docker-scan && \
    chmod +x $DOCKER_CONFIG/cli-plugins/docker-scan
RUN curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
RUN curl -fsSL https://install-cli.jfrog.io | sh
COPY --from=helm-docs /usr/bin/helm-docs /usr/bin/helm-docs
COPY bin/linux-amd64/winch /usr/local/bin/winch
COPY docker/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
