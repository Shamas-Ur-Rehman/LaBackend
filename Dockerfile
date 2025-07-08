# Dockerfile
FROM golang:1.21

# Install TLS root certs for MongoDB Atlas
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
