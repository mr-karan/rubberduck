FROM golang:1.12-alpine AS backend-build
WORKDIR /rubberduck/
COPY ./ ./
RUN go build -o rubberduck
# ---- Release Alpine Image ----
FROM alpine:latest AS deploy
RUN set -ex \
    && echo "http://nl.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache \
    iftop \
    drill \
    strace \
    curl \
    iperf \
    nmap \
    conntrack-tools \
    bash \
    jq \
    vim \
    nano \
    tree \
    ca-certificates \
    bind-tools
WORKDIR /rubberduck/
COPY --from=backend-build /rubberduck/rubberduck ./
CMD [ "bash", "-c", "./rubberduck" ]
