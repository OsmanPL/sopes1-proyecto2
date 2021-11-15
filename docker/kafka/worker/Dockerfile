FROM golang:1.17.1-alpine3.14

RUN mkdir -p /api/go

WORKDIR /api/go

COPY . .

RUN go mod download
RUN go mod vendor

CMD ["go", "run", "worker/worker.go"]