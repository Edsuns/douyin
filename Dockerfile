FROM golang:1.18-alpine

WORKDIR $GOPATH/douyin
COPY . .

RUN apk add --update ffmpeg~=4.4.1-r2

RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go build -o ./bin/app ./app

EXPOSE 8080
ENTRYPOINT ["./bin/app"]
