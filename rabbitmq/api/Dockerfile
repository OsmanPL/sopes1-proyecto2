FROM golang:1.17.1-alpine3.14

RUN mkdir -p /api/go

WORKDIR /api/go

COPY . .

RUN go mod download

CMD ["go", "run", "worker/worker.go"]