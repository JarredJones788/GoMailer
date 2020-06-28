FROM alpine:3.8

WORKDIR /app

COPY ./server ./server

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

EXPOSE 5000
CMD ["/app/server", "-production"]