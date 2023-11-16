FROM golang:1.21-alpine

# Using apk package manager for Alpine Linux
RUN apk update && apk add --no-cache git

WORKDIR /app

RUN git clone https://github.com/akshay0074700747/e-commerce.git .

RUN go mod download

WORKDIR /app/cmd

COPY ../cmd/.env /app/cmd/.env

RUN go build -o bin/ecommerce-executable

WORKDIR /app/cmd/bin

CMD ["./ecommerce-executable"]