FROM golang:1.17 as builder
WORKDIR /usr/src/app
ENV GOPROXY https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o http-server

FROM alpine as runner
WORKDIR /opt/app
COPY --from=builder /usr/src/app/http-server /opt/app/
EXPOSE 8080
ENTRYPOINT ["http-server"]