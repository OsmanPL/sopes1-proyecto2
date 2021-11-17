FROM golang:1.17.1-alpine3.14

RUN mkdir -p /api/go-client-grpc

WORKDIR /api/go-client-grpc

COPY . .

RUN go mod download

EXPOSE 4000

CMD ["go", "run", "client/client.go"]