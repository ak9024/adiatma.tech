FROM golang:alpine AS builder

ARG PORT
ARG HOST
ARG BASE_PATH

ENV PORT=${PORT}
ARG HOST=${HOST}
ARG BASE_PATH=${BASE_PATH}

WORKDIR /app/api

COPY go.mod go.sum ./

COPY . .

RUN go run . env

RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM alpine:latest

WORKDIR /app/api

COPY --from=builder /app/api .
COPY --from=builder /app/api/.env .

ENTRYPOINT [ "./api" ]
