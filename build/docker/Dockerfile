FROM golang:1.21.12-alpine3.20 AS builder
WORKDIR /server

COPY . .
RUN apk add build-base
RUN go build -o main

FROM alpine:latest
ARG BUILD_TIME
ARG SHA1_VER

RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl tzdata && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /server
COPY --from=builder /server/main /server/main
COPY --from=builder /server/configs/app-docker.yml /server/app.yml

RUN ls
ENV SHA1_VER=${SHA1_VER}
ENV BUILD_TIME=${BUILD_TIME}
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /server/main
USER appuser

ENTRYPOINT [ "./main" ]
CMD [ "jubobe" ]
