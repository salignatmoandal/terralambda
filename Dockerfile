FROM golang:1.21-alpine

WORKDIR /app

RUN apk add --no-cache \
    terraform \
    zip \
    git \
    aws-cli

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/terralambda ./cmd/main.go

ENTRYPOINT ["terralambda"]