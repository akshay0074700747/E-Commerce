# Build stage
FROM golang:1.21-bullseye AS build

RUN apt-get update && apt-get install -y git

WORKDIR /app

RUN git clone https://github.com/akshay0074700747/e-commerce.git .

RUN go mod download

WORKDIR /app/cmd

COPY ../cmd/.env /app/cmd/bin/.env

RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o bin/ecommerce-executable

# Final stage
FROM debian:bullseye-slim

ENV GIN_MODE release

WORKDIR /ecommerce-executable


COPY --from=build /app/cmd/bin/ecommerce-executable .


COPY --from=build /app/cmd/bin/.env .


COPY --from=build /app/cmd/payment.html .

# EXPOSE 3000

CMD ["./ecommerce-executable"]
