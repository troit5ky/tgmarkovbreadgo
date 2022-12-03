FROM golang:1.19.3-alpine3.16

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN apk add g++ && \
    go build -v -o /usr/local/bin/app main.go

CMD ["app"]