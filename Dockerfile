FROM golang:1.23.1

# Install CA certs for MongoDB TLS
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
