FROM golang:alpine AS builder

ARG PORT
ENV PORT=${PORT}

WORKDIR /app/api

COPY go.mod go.sum ./

COPY . .

RUN go run . env

RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM alpine:latest

WORKDIR /app/api

COPY --from=builder /app/api .
COPY --from=builder /app/api/.env .

ENTRYPOINT [ "./api" ]
