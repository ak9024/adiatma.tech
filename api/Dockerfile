FROM golang:alpine AS builder

WORKDIR /app/api

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM alpine:latest

WORKDIR /app/api

COPY --from=builder /app/api .

ENTRYPOINT [ "./api" ]
