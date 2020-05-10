FROM golang:1.14.1

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN echo $GOPATH

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build homepage-service

EXPOSE 8080
ENTRYPOINT ["./homepage-service"]
