FROM golang:1.17.1-alpine3.14

RUN mkdir -p /api/go-server-grpc

WORKDIR /api/go-server-grpc

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "server/server.go"]