FROM golang:1.13.6-alpine3.11 AS stage
ADD . /build
RUN cd /build && go build -v -o gomail src/*.go \
    && mkdir dist && cat .dist | xargs -I % cp -R % dist

FROM alpine:3
RUN apk --no-cache add curl figlet
WORKDIR /app
COPY --from=stage /build/dist/ /app/
EXPOSE 58725
ENTRYPOINT ["/bin/sh", "-c", "echo '\n🌐 '$(curl -s ipecho.net/plain) && ./gomail"]
