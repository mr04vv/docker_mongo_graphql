FROM golang:1.12.0-alpine3.9

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git 

RUN go get github.com/mr04vv/kamimai/cmd/kamimai
# freshのインストール
RUN go get github.com/pilu/fresh
# イメージからコンテナを作成する際に実行
EXPOSE 1323
CMD ["fresh"]

