FROM golang:1.18-alpine

WORKDIR $GOPATH/douyin
COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --update ffmpeg~=4.4.1-r2

RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go build -o ./douyin ./app

EXPOSE 8080
CMD ["./douyin"]
