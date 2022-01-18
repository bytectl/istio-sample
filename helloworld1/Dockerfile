FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct  \
    GO111MODULE=on \
    GONOSUMDB="github.com" \
    make build

FROM wuxingzhong/distroless:static

VOLUME /data/configs

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /data/configs
WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./helloworld1", "-conf", "/data/configs"]
